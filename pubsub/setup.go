package pubsub

import (
	"fmt"
	"log"
)

func (br *PubSubBenchmarkRun) Setup() error {
	queries := []string{

		"DROP TABLE IF EXISTS consumer_offsets CASCADE;",
		"DROP TABLE IF EXISTS log_counter CASCADE;",
		"DROP TABLE IF EXISTS topicpartition CASCADE;",

		// the core log. it's as if it's one partition
		`
CREATE TABLE topicpartition (
  id          BIGSERIAL PRIMARY KEY,  -- physical row id (not used for reads)
  c_offset      BIGINT UNIQUE NOT NULL, -- gapless, strictly increasing, read-key
  payload     BYTEA NOT NULL,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
		`,

		// fast range scans on the log via offset
		"CREATE UNIQUE INDEX idx_topicpartition_offset ON topicpartition(offset);",

		// a single table acting as a global counter for writers. there's only one row they increase
		// (ofc this can be extended for more tables)
		`
CREATE TABLE log_counter (
 id           INT PRIMARY KEY CHECK (id = 1),
 next_offset  BIGINT NOT NULL
);`,

		// initialize the counter
		"INSERT INTO log_counter(id, next_offset) VALUES (1, 1);",

		// The consumer group offsets table. Each group has their own row they contend for
		`
CREATE TABLE consumer_offsets (
  group_id     TEXT PRIMARY KEY,
  next_offset  BIGINT NOT NULL DEFAULT 1 -- next offset to claim/read
);`,
	}

	for _, q := range queries {
		if _, err := br.Db.ExecContext(br.Ctx, q); err != nil {
			return fmt.Errorf("setup: %w", err)
		}
	}

	log.Println("[pub info] successfully dropped + applied pub-sub tables")
	return nil
}
