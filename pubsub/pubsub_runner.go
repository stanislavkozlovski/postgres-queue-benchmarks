package pubsub

import (
	"context"
	"database/sql"
	"golang.org/x/time/rate"
	c "main/common"
	"sync"
	"time"
)

type PubSubConfig struct {
	c.BaselineConfig

	// Readers denotes the number of subscribers PER group
	NumConsumerGroups int

	// the max number of messages to read per request
	// in Kafka, the defaults are `fetch.max.bytes` (50MB) and `max.partition.fetch.bytes` (1MB)
	ReadBatchSize int

	// the number of messages per read/write request. In Kafka, linger.ms and batch.size (16 KB) control this per partition.
	// Its 16 KiB per producer per partition by default too, but increasing it is recommended in general
	WriteBatchSize int

	// Denotes whether readers will use at-least-once and strict ordering.
	// If using this, opt for larger read batches and only one consumer per group, as we're serializing reads on the log per group.
	// Obviously there are no partitions here to shard the data, so we're limited in what one process can read.
	KafkaSemantics bool
}

type PubSubBenchmarkRun struct {
	config *PubSubConfig
	*c.BenchmarkRun
	PubSubMetrics *PubSubMetrics
}

func NewPubSubBenchmarkRun(cfg *PubSubConfig, db *sql.DB, ctx context.Context, writeLimiter *rate.Limiter) (*PubSubBenchmarkRun, error) {
	metrics := NewPubSubMetrics(cfg.NumConsumerGroups, cfg.Readers)
	return &PubSubBenchmarkRun{
		config: cfg,
		BenchmarkRun: c.NewBenchmarkRun(db,
			c.NewMetrics(cfg.Writers, cfg.Readers), ctx, writeLimiter),
		PubSubMetrics: metrics,
	}, nil
}

func (br *PubSubBenchmarkRun) Run() {
	var wg sync.WaitGroup
	for i := 0; i < br.config.Writers; i++ {
		wg.Add(1)
		go br.Writer(i, &wg)
	}

	for groupID := 0; groupID < br.config.NumConsumerGroups; groupID++ {
		for consumerID := 0; consumerID < br.config.Readers; consumerID++ {
			wg.Add(1)
			// groupID int, gm *GroupMetrics, consumerID int, wg *sync.WaitGroup,
			//	kafkaSemantics bool
			go br.GroupMember(groupID, br.PubSubMetrics.Groups[groupID], consumerID, &wg, br.config.KafkaSemantics)
		}
	}
	wg.Add(1)
	go br.Reporter(&wg)

	<-br.Ctx.Done()
	time.Sleep(50 * time.Millisecond)
	wg.Wait()
}
