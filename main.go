package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"time"

	c "main/common"
	"main/queue"
)

func main() {
	var (
		host   = flag.String("host", "localhost", "PostgreSQL host")
		port   = flag.Int("port", 5432, "PostgreSQL port")
		dbName = flag.String("db", "benchmark", "Database name")
		user   = flag.String("user", "postgres", "Database user")
		pass   = flag.String("password", "", "Database password")

		writers        = flag.Int("writers", 4, "Number of writers")
		readers        = flag.Int("readers", 4, "Number of readers")
		duration       = flag.Duration("duration", 30*time.Second, "Test duration")
		payload        = flag.Int("payload", 1024, "Payload size in bytes")
		reportEvery    = flag.Duration("report", 5*time.Second, "Report interval")
		throttleWrites = flag.Int("throttle_writes", 0, "Throttle writer rows/sec (0=unlimited)")
		tuneTableVac   = flag.Bool("tune-table-vacuum", false, "Apply aggressive autovacuum/fillfactor to queue table")
		mode           = flag.String("mode", "queue", "mode: queue|pubsub - which benchmark to run")
	)
	flag.Parse()

	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		*host, *port, *user, *dbName)
	if *pass != "" {
		connStr += " password=" + *pass
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("connect to postgres: %w", err)
	}

	throttleWrs := *throttleWrites
	var limiter *rate.Limiter
	if throttleWrs > 0 {
		limiter = rate.NewLimiter(rate.Limit(throttleWrs), throttleWrs)
	}

	baseCfg := &c.BaselineConfig{
		Writers:        *writers,
		Readers:        *readers,
		Duration:       *duration,
		PayloadSize:    *payload,
		ReportInterval: *reportEvery,
	}

	ctx, _ := context.WithTimeout(context.Background(), baseCfg.Duration)

	if *mode == "queue" {
		queueCfg := &queue.QueueConfig{
			BaselineConfig:  *baseCfg,
			TuneTableVacuum: *tuneTableVac,
		}

		br, err := queue.NewQueueBenchmarkRun(queueCfg, db, ctx, limiter)
		if err != nil {
			log.Fatalf("connect: %v", err)
		}
		if err := br.Setup(); err != nil {
			log.Fatalf("setup: %v", err)
		}
		br.Run()
		br.PrintSummary(queueCfg.Duration)
		_ = br.Db.Close()
		log.Println("queue benchmark complete")
	}
}
