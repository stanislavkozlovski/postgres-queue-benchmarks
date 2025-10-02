package pubsub

import (
	"fmt"
	"time"
)

// PrintSummary prints a per-group summary of pubsub metrics.
// dur = benchmark duration; kafkaSemantics controls how we label delivery semantics.
func (pm *PubSubMetrics) PrintSummary(dur time.Duration, kafkaSemantics bool) {
	fmt.Printf("\n=== PubSub Summary ===\n")
	secs := dur.Seconds()

	semantics := "at-most-once (best effort)"
	if kafkaSemantics {
		semantics = "at-least-once (strict ordering)"
	}
	fmt.Printf("Delivery Semantics: %s\n", semantics)
	fmt.Printf("Duration: %s (%.1fs)\n", dur, secs)

	snapshots := pm.Snapshot()
	for _, gs := range snapshots {
		fmt.Printf("\n-- Consumer Group %d --\n", gs.GroupID)
		fmt.Printf("  Reads Completed:   %d\n", gs.ReadsCompleted)
		fmt.Printf("  Read Errors:       %d\n", gs.ReadErrors)
		fmt.Printf("  Updates Completed: %d\n", gs.UpdatesCompleted)
		fmt.Printf("  Claim Errors:      %d\n", gs.ClaimErrors)
		fmt.Printf("  Empty Claims:      %d\n", gs.EmptyClaims)
		fmt.Printf("  Avg Poll Size:     %.2f\n", gs.AvgPollSize)

		if secs > 0 {
			fmt.Printf("  Avg Throughput:    %.2f reads/sec\n", float64(gs.ReadsCompleted)/secs)
		}

		fmt.Printf("\n  Read Latency (claim txn):\n")
		fmt.Printf("    P50: %v\n", time.Duration(gs.ReadP50))
		fmt.Printf("    P95: %v\n", time.Duration(gs.ReadP95))
		fmt.Printf("    P99: %v\n", time.Duration(gs.ReadP99))

		fmt.Printf("\n  End-to-End Latency (created_at → consumed):\n")
		fmt.Printf("    P50: %v\n", time.Duration(gs.E2EP50))
		fmt.Printf("    P95: %v\n", time.Duration(gs.E2EP95))
		fmt.Printf("    P99: %v\n", time.Duration(gs.E2EP99))
	}
	fmt.Println()
}
