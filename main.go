package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/time/rate"

	_ "github.com/lib/pq"
)

type Config struct {
	ConnStr        string
	Writers        int
	Readers        int
	Duration       time.Duration
	PayloadSize    int
	ReportInterval time.Duration
	ThrottleWrites int // rows/sec, 0 = unlimited
}

type BenchmarkRun struct {
	config  *Config
	db      *sql.DB
	metrics *Metrics
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewBenchmarkRun(cfg *Config) (*BenchmarkRun, error) {
	db, err := sql.Open("postgres", cfg.ConnStr)
	if err != nil {
		return nil, fmt.Errorf("connect: %w", err)
	}
	target := cfg.Writers + cfg.Readers
	db.SetMaxOpenConns(target)
	db.SetMaxIdleConns(target)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Duration)

	return &BenchmarkRun{
		config:  cfg,
		db:      db,
		metrics: NewMetrics(cfg.Writers, cfg.Readers),
		ctx:     ctx,
		cancel:  cancel,
	}, nil
}

func (br *BenchmarkRun) Setup() error {
	queries := []string{
		"DROP TABLE IF EXISTS queue_archive",
		"DROP TABLE IF EXISTS queue",
		`CREATE TABLE queue (
			id BIGSERIAL PRIMARY KEY,
			payload BYTEA NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`,
		`CREATE TABLE queue_archive (
			id BIGINT,
			payload BYTEA NOT NULL,
			created_at TIMESTAMP NOT NULL,
			read_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`,
		"CREATE INDEX idx_queue_id ON queue(id)",
	}
	for _, q := range queries {
		if _, err := br.db.ExecContext(br.ctx, q); err != nil {
			return fmt.Errorf("setup: %w", err)
		}
	}
	return nil
}

func (br *BenchmarkRun) Writer(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, _ := br.db.Conn(br.ctx)
	defer conn.Close()

	hist := br.metrics.writerHists[id]
	payload := make([]byte, br.config.PayloadSize)
	_, _ = rand.Read(payload)

	stmt, _ := conn.PrepareContext(br.ctx, "INSERT INTO queue (payload) VALUES ($1)")
	defer stmt.Close()

	// set up rate limiter if configured
	var limiter *rate.Limiter
	if br.config.ThrottleWrites > 0 {
		limiter = rate.NewLimiter(rate.Limit(br.config.ThrottleWrites), br.config.ThrottleWrites)
	}

	for {
		select {
		case <-br.ctx.Done():
			return
		default:
			// throttle if limiter active
			if limiter != nil {
				if err := limiter.Wait(br.ctx); err != nil {
					return
				}
			}

			start := time.Now()
			if _, err := stmt.ExecContext(br.ctx, payload); err != nil {
				br.metrics.WriteErrors.Add(1)
			} else {
				br.metrics.WritesCompleted.Add(1)
			}
			lat := time.Since(start).Nanoseconds()
			if err := hist.RecordValue(lat); err != nil {
				log.Printf("[writer %d] histogram record error: %v (lat=%dns)", id, err, lat)
			}
		}
	}
}

func (br *BenchmarkRun) Reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, _ := br.db.Conn(br.ctx)
	defer conn.Close()

	hist := br.metrics.readerHists[id]

	const selSQL = `
		SELECT id, payload, created_at
		FROM queue
		ORDER BY id
		FOR UPDATE SKIP LOCKED
		LIMIT 1
	`
	selStmt, _ := conn.PrepareContext(br.ctx, selSQL)
	defer selStmt.Close()

	delStmt, _ := conn.PrepareContext(br.ctx,
		"DELETE FROM queue WHERE id = $1")
	defer delStmt.Close()

	insStmt, _ := conn.PrepareContext(br.ctx,
		"INSERT INTO queue_archive (id, payload, created_at, read_at) VALUES ($1,$2,$3,NOW())")
	defer insStmt.Close()

	for {
		select {
		case <-br.ctx.Done():
			return
		default:
			start := time.Now()
			tx, err := conn.BeginTx(br.ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
			if err != nil {
				br.metrics.ReadErrors.Add(1)
				continue
			}

			var (
				id64    int64
				payload []byte
				created time.Time
			)
			// step 1: read + claim the row
			if err := tx.Stmt(selStmt).QueryRowContext(br.ctx).Scan(&id64, &payload, &created); err != nil {
				_ = tx.Rollback()
				time.Sleep(200 * time.Microsecond)
				continue
			}

			// step 2: process
			// (simulate some work — currently just no-op)
			_ = payload // you could hash, sleep, etc.

			// step 3a: delete
			if _, err := tx.Stmt(delStmt).ExecContext(br.ctx, id64); err != nil {
				_ = tx.Rollback()
				br.metrics.ReadErrors.Add(1)
				continue
			}
			// step 3b: insert into archive
			if _, err := tx.Stmt(insStmt).ExecContext(br.ctx, id64, payload, created); err != nil {
				_ = tx.Rollback()
				br.metrics.ReadErrors.Add(1)
				continue
			}

			if err := tx.Commit(); err != nil {
				br.metrics.ReadErrors.Add(1)
				continue
			}

			br.metrics.ReadsCompleted.Add(1)
			br.metrics.UpdatesCompleted.Add(1)
			lat := time.Since(start).Nanoseconds()
			if err := hist.RecordValue(lat); err != nil {
				log.Printf("[reader %d] histogram record error: %v (lat=%dns)", id, err, lat)
			}
		}
	}
}

func (br *BenchmarkRun) Run() {
	var wg sync.WaitGroup
	for i := 0; i < br.config.Writers; i++ {
		wg.Add(1)
		go br.Writer(i, &wg)
	}
	for i := 0; i < br.config.Readers; i++ {
		wg.Add(1)
		go br.Reader(i, &wg)
	}
	wg.Add(1)
	go br.Reporter(&wg)

	<-br.ctx.Done()
	time.Sleep(50 * time.Millisecond)
	wg.Wait()
}

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
	)
	flag.Parse()

	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		*host, *port, *user, *dbName)
	if *pass != "" {
		connStr += " password=" + *pass
	}

	cfg := &Config{
		ConnStr:        connStr,
		Writers:        *writers,
		Readers:        *readers,
		Duration:       *duration,
		PayloadSize:    *payload,
		ReportInterval: *reportEvery,
		ThrottleWrites: *throttleWrites,
	}

	br, err := NewBenchmarkRun(cfg)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	if err := br.Setup(); err != nil {
		log.Fatalf("setup: %v", err)
	}
	br.Run()
	br.PrintSummary()
	_ = br.db.Close()
	log.Println("benchmark complete")
}
