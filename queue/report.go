package queue

import (
	"fmt"
	"sync"
	"time"
)

func (br *QueueBenchmarkRun) Reporter(wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(br.config.ReportInterval)
	defer ticker.Stop()

	lastWrites := int64(0)
	lastReads := int64(0)
	lastTime := time.Now()

	for {
		select {
		case <-br.Ctx.Done():
			return
		case <-ticker.C:
			now := time.Now()
			secs := now.Sub(lastTime).Seconds()

			writes := br.Metrics.WritesCompleted.Load()
			reads := br.Metrics.ReadsCompleted.Load()

			writeRate := float64(writes-lastWrites) / secs
			readRate := float64(reads-lastReads) / secs
			queueDepth := writes - reads

			fmt.Printf("[%s] W: %.0f/s R: %.0f/s QDepth: %d Err(W/R): %d/%d\n",
				now.Format("15:04:05"),
				writeRate,
				readRate,
				queueDepth,
				br.Metrics.WriteErrors.Load(),
				br.Metrics.ReadErrors.Load(),
			)

			lastWrites = writes
			lastReads = reads
			lastTime = now
		}
	}
}
