package main

import (
	"fmt"
	"github.com/HdrHistogram/hdrhistogram-go"
	"log"
	"sync"
	"time"
)

func (br *BenchmarkRun) Reporter(wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(br.config.ReportInterval)
	defer ticker.Stop()

	lastWrites := int64(0)
	lastReads := int64(0)
	lastTime := time.Now()

	for {
		select {
		case <-br.ctx.Done():
			return
		case <-ticker.C:
			now := time.Now()
			secs := now.Sub(lastTime).Seconds()

			writes := br.metrics.WritesCompleted.Load()
			reads := br.metrics.ReadsCompleted.Load()

			writeRate := float64(writes-lastWrites) / secs
			readRate := float64(reads-lastReads) / secs
			queueDepth := writes - reads

			fmt.Printf("[%s] W: %.0f/s R: %.0f/s QDepth: %d Err(W/R): %d/%d\n",
				now.Format("15:04:05"),
				writeRate,
				readRate,
				queueDepth,
				br.metrics.WriteErrors.Load(),
				br.metrics.ReadErrors.Load(),
			)

			lastWrites = writes
			lastReads = reads
			lastTime = now
		}
	}
}

func mergeHists(hists []*hdrhistogram.Histogram) *hdrhistogram.Histogram {
	if len(hists) == 0 {
		return nil
	}
	// 1ns to 1m range, 3 sig figs
	out := hdrhistogram.New(1, int64(time.Minute), 3)

	var totalDropped int64
	for _, h := range hists {
		if h == nil {
			continue
		}
		dropped := out.Merge(h)
		totalDropped += dropped
	}
	if totalDropped > 0 {
		log.Printf("[merge] dropped %d values outside histogram range", totalDropped)
	}
	return out
}
func (br *BenchmarkRun) PrintSummary() {
	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Total Writes: %d\n", br.metrics.WritesCompleted.Load())
	fmt.Printf("Total Reads: %d\n", br.metrics.ReadsCompleted.Load())
	fmt.Printf("Total Updates: %d\n", br.metrics.UpdatesCompleted.Load())
	fmt.Printf("Write Errors: %d\n", br.metrics.WriteErrors.Load())
	fmt.Printf("Read Errors: %d\n", br.metrics.ReadErrors.Load())

	secs := br.config.Duration.Seconds()
	if secs > 0 {
		fmt.Printf("Avg Write Throughput: %.2f rows/sec\n", float64(br.metrics.WritesCompleted.Load())/secs)
		fmt.Printf("Avg Read Throughput: %.2f rows/sec\n", float64(br.metrics.ReadsCompleted.Load())/secs)
	}

	writeHist := mergeHists(br.metrics.writerHists)
	readHist := mergeHists(br.metrics.readerReadHists)
	e2eHist := mergeHists(br.metrics.readerE2EHists)

	if writeHist != nil {
		fmt.Printf("\nWrite Latencies (INSERT only):\n")
		fmt.Printf("  P50: %v\n", time.Duration(writeHist.ValueAtQuantile(50)))
		fmt.Printf("  P95: %v\n", time.Duration(writeHist.ValueAtQuantile(95)))
		fmt.Printf("  P99: %v\n", time.Duration(writeHist.ValueAtQuantile(99)))
	}
	if readHist != nil {
		fmt.Printf("\nRead Latencies (txn: SELECT+DELETE+INSERT):\n")
		fmt.Printf("  P50: %v\n", time.Duration(readHist.ValueAtQuantile(50)))
		fmt.Printf("  P95: %v\n", time.Duration(readHist.ValueAtQuantile(95)))
		fmt.Printf("  P99: %v\n", time.Duration(readHist.ValueAtQuantile(99)))
	}
	if e2eHist != nil {
		fmt.Printf("\nEnd-to-End Latencies (created_at → consumed):\n")
		fmt.Printf("  P50: %v\n", time.Duration(e2eHist.ValueAtQuantile(50)))
		fmt.Printf("  P95: %v\n", time.Duration(e2eHist.ValueAtQuantile(95)))
		fmt.Printf("  P99: %v\n", time.Duration(e2eHist.ValueAtQuantile(99)))
	}
	fmt.Println()
}
