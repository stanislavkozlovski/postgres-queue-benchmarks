package pubsub

import (
	"database/sql"
	"log"
	"sync"
	"time"
)

// consumers get the latest offset via the same log_counter table
// then atomically claim a range of messages via consumer_offsets
// This has at least once semantics
// It returns the start/end offset that the consumer just claimed.
const claimSQL = `
WITH counter_tip AS (
  SELECT (next_offset - 1) AS highest_committed_offset
  FROM log_counter
  WHERE id = 1
),

claimed_range AS (
  UPDATE consumer_offsets c
  SET next_offset = c.next_offset
                  + LEAST(
                      $2,
                      GREATEST(0,
                        (SELECT highest_committed_offset FROM counter_tip) - c.next_offset + 1
                      )
                    )
  WHERE c.group_id = $1
  RETURNING
    (next_offset - LEAST(
       $2,
       GREATEST(0, (SELECT highest_committed_offset FROM counter_tip) - next_offset + $2)
     )) AS claimed_start_offset,
    (next_offset - 1) AS claimed_end_offset
)
SELECT claimed_start_offset, claimed_end_offset
FROM claimed_range;
`

const readSQL = `SELECT c_offset, payload, created_at
		   	FROM topicpartition
		  	WHERE c_offset BETWEEN $1 AND $2
		  	ORDER BY c_offset`

// GroupMember runs the read loop for a single subscriber in a consumer group.
// groupID is the unique group identifier.
// kafkaSemantics toggles whether to use At-Least-Once and Strict Order in processing;
func (br *PubSubBenchmarkRun) GroupMember(groupID int, gm *GroupMetrics, consumerID int, wg *sync.WaitGroup,
	kafkaSemantics bool) {
	defer wg.Done()

	conn, _ := br.Db.Conn(br.Ctx)
	defer conn.Close()

	claimStmt, err := conn.PrepareContext(br.Ctx, claimSQL)
	if err != nil {
		panic("prepare claimStmt failed: " + err.Error())
	}
	defer claimStmt.Close()

	readStmt, err := conn.PrepareContext(br.Ctx, readSQL)
	if err != nil {
		panic("prepare readStmt failed: " + err.Error())
	}
	defer readStmt.Close()

	for {
		select {
		case <-br.Ctx.Done():
			return
		default:
			if kafkaSemantics {
				br.kafkaSemanticRead(conn, gm, groupID, consumerID, claimStmt, readStmt)
			} else {
				br.atMostOnceRead(conn, gm, groupID, consumerID, claimStmt, readStmt)
			}
		}
	}
}

// atMostOnce does the Claim Offset in a transaction, and then reads/processes the data outside it.
// This leads to at most once semantics, as claims can be lost if the reader crashes after claiming and doesn't process.
// Retries could be added to ensure at-least-once, but that's extra work.
// This also breaks ordering guarantees, because during errors/retries/timeouts other consumers can claim and process future messages before this one.
func (br *PubSubBenchmarkRun) atMostOnceRead(conn *sql.Conn, gm *GroupMetrics, groupID int, consumerID int, claimOffsetsStmt *sql.Stmt, readDataStmt *sql.Stmt) {
	// -- Claim TX Start
	claimOffsetTx, err := conn.BeginTx(br.Ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		gm.ReadErrors.Add(1)
		return
	}

	var startOff, endOff sql.NullInt64
	if err := claimOffsetTx.StmtContext(br.Ctx, claimOffsetsStmt).QueryRowContext(br.Ctx,
		groupID, br.config.ReadBatchSize).Scan(&startOff, &endOff); err != nil {
		_ = claimOffsetTx.Rollback()
		log.Printf("[consumer g%d r%d] Claim err: %v", groupID, consumerID, err)

		gm.ClaimErrors.Add(1)
		time.Sleep(jitter(100*time.Microsecond, 800000*time.Microsecond))
		return
	}

	if !startOff.Valid || !endOff.Valid || startOff.Int64 > endOff.Int64 {
		// nothing to claim
		_ = claimOffsetTx.Commit()
		gm.EmptyClaims.Add(1)
		time.Sleep(jitter(100*time.Microsecond, 2*time.Millisecond))
		return
	}

	if err := claimOffsetTx.Commit(); err != nil {
		gm.ClaimErrors.Add(1)
		return
	}
	// -- TX is over

	// read the claimed range
	start := time.Now()
	rows, err := readDataStmt.QueryContext(br.Ctx, startOff.Int64, endOff.Int64)
	if err != nil {
		gm.ReadErrors.Add(1)
		return
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		var off int64
		var payload []byte
		var created time.Time
		if err := rows.Scan(&off, &payload, &created); err != nil {
			log.Printf("[consumer g%d r%d] scan err: %v", groupID, consumerID, err)
			continue
		}
		count++

		// TODO: “process payload”

		e2eLat := time.Since(created).Nanoseconds()
		_ = gm.ReaderE2EHists[consumerID].RecordValue(e2eLat)
	}

	// record metrics
	gm.ReadsCompleted.Add(count)
	gm.UpdatesCompleted.Add(count)
	gm.PolledRecords.Add(count)
	gm.Polls.Add(1)
	lat := time.Since(start).Nanoseconds()
	_ = gm.ReaderReadHists[consumerID].RecordValue(lat)
}

// kafkaSemanticRead does the Offset Claim + Select in one transaction.
// This gives you at-least-once semantics and strict ordering.
// It doesn't make much sense to run a lot of consumers in a group with this mode.
func (br *PubSubBenchmarkRun) kafkaSemanticRead(conn *sql.Conn, gm *GroupMetrics, groupID int, consumerID int, claimOffsetsStmt *sql.Stmt, readDataStmt *sql.Stmt) {
	tx, err := conn.BeginTx(br.Ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		gm.ReadErrors.Add(1)
		return
	}

	var startOff, endOff sql.NullInt64
	if err := tx.StmtContext(br.Ctx, claimOffsetsStmt).QueryRowContext(br.Ctx,
		groupID, br.config.ReadBatchSize).Scan(&startOff, &endOff); err != nil {
		_ = tx.Rollback()
		gm.ClaimErrors.Add(1)
		time.Sleep(jitter(100*time.Microsecond, 800*time.Microsecond))
		return
	}

	if !startOff.Valid || !endOff.Valid || startOff.Int64 > endOff.Int64 {
		// nothing to claim
		_ = tx.Commit()
		gm.EmptyClaims.Add(1)
		time.Sleep(jitter(100*time.Microsecond, 2*time.Millisecond))
		return
	}

	// read the claimed range
	start := time.Now()
	rows, err := tx.StmtContext(br.Ctx, readDataStmt).QueryContext(br.Ctx, startOff.Int64, endOff.Int64)
	if err != nil {
		_ = tx.Rollback()
		gm.ReadErrors.Add(1)
		return
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		var off int64
		var payload []byte
		var created time.Time
		if err := rows.Scan(&off, &payload, &created); err != nil {
			log.Printf("[consumer g%d r%d] scan err: %v", groupID, consumerID, err)
			continue
		}
		count++

		// TODO: “process payload”

		e2eLat := time.Since(created).Nanoseconds()
		_ = gm.ReaderE2EHists[consumerID].RecordValue(e2eLat)
	}

	if err := tx.Commit(); err != nil {
		gm.ReadErrors.Add(1)
		return
	}

	// record metrics
	gm.ReadsCompleted.Add(count)
	gm.UpdatesCompleted.Add(count)
	gm.PolledRecords.Add(count)
	gm.Polls.Add(1)
	lat := time.Since(start).Nanoseconds()
	_ = gm.ReaderReadHists[consumerID].RecordValue(lat)
}

func jitter(min, max time.Duration) time.Duration {
	if max <= min {
		return min
	}
	// fast xorshift
	r := time.Now().UnixNano()
	r ^= r << 13
	r ^= r >> 7
	r ^= r << 17
	span := int64(max - min)
	if span <= 0 {
		return min
	}
	return min + time.Duration(r%span)
}
