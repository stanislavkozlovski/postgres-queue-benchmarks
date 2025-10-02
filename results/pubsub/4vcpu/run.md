```bash
=== Summary ===
Total Writes: 2995200
Total Reads: 0
Total Updates: 0
Write Errors: 100
Read Errors: 0
Avg Write Throughput: 24960.00 rows/sec
Avg Read Throughput: 0.00 rows/sec

Write Latencies (INSERT only):
  P50: 4.804607ms
  P95: 6.045695ms
  P99: 198.574079ms

=== PubSub Summary ===
Delivery Semantics: at-least-once (strict ordering)
Duration: 2m0s (120.0s)

-- Consumer Group 0 --
  Reads Completed:   2994900
  Read Errors:       0
  Updates Completed: 2994900
  Claim Errors:      1
  Empty Claims:      895
  Avg Poll Size:     146.59
  Avg Throughput:    24957.50 reads/sec

  Read Latency (claim txn):
    P50: 6.832127ms
    P95: 7.794687ms
    P99: 8.339455ms

  End-to-End Latency (created_at → consumed):
    P50: 11.567103ms
    P95: 17.661951ms
    P99: 205.783039ms

-- Consumer Group 1 --
  Reads Completed:   2994800
  Read Errors:       1
  Updates Completed: 2994800
  Claim Errors:      0
  Empty Claims:      1497
  Avg Poll Size:     141.04
  Avg Throughput:    24956.67 reads/sec

  Read Latency (claim txn):
    P50: 6.430719ms
    P95: 7.319551ms
    P99: 7.880703ms

  End-to-End Latency (created_at → consumed):
    P50: 11.321343ms
    P95: 16.973823ms
    P99: 206.700543ms

-- Consumer Group 2 --
  Reads Completed:   2994800
  Read Errors:       1
  Updates Completed: 2994800
  Claim Errors:      1
  Empty Claims:      1728
  Avg Poll Size:     138.57
  Avg Throughput:    24956.67 reads/sec

  Read Latency (claim txn):
    P50: 6.275071ms
    P95: 7.127039ms
    P99: 7.700479ms

  End-to-End Latency (created_at → consumed):
    P50: 11.149311ms
    P95: 16.752639ms
    P99: 207.355903ms

-- Consumer Group 3 --
  Reads Completed:   2994900
  Read Errors:       0
  Updates Completed: 2994900
  Claim Errors:      1
  Empty Claims:      243
  Avg Poll Size:     133.78
  Avg Throughput:    24957.50 reads/sec

  Read Latency (claim txn):
    P50: 6.619135ms
    P95: 7.368703ms
    P99: 7.921663ms

  End-to-End Latency (created_at → consumed):
    P50: 10.870783ms
    P95: 14.835711ms
    P99: 206.700543ms

-- Consumer Group 4 --
  Reads Completed:   2994900
  Read Errors:       0
  Updates Completed: 2994900
  Claim Errors:      1
  Empty Claims:      1815
  Avg Poll Size:     140.22
  Avg Throughput:    24957.50 reads/sec

  Read Latency (claim txn):
    P50: 6.311935ms
    P95: 7.180287ms
    P99: 7.692287ms

  End-to-End Latency (created_at → consumed):
    P50: 11.263999ms
    P95: 16.826367ms
    P99: 206.700543ms

-- Consumer Group 5 --
  Reads Completed:   2994800
  Read Errors:       1
  Updates Completed: 2994800
  Claim Errors:      1
  Empty Claims:      532
  Avg Poll Size:     143.48
  Avg Throughput:    24956.67 reads/sec

  Read Latency (claim txn):
    P50: 6.860799ms
    P95: 7.737343ms
    P99: 8.278015ms

  End-to-End Latency (created_at → consumed):
    P50: 11.395071ms
    P95: 17.022975ms
    P99: 207.355903ms

-- Consumer Group 6 --
  Reads Completed:   2994800
  Read Errors:       1
  Updates Completed: 2994800
  Claim Errors:      1
  Empty Claims:      2173
  Avg Poll Size:     155.06
  Avg Throughput:    24956.67 reads/sec

  Read Latency (claim txn):
    P50: 6.758399ms
    P95: 7.811071ms
    P99: 8.343551ms

  End-to-End Latency (created_at → consumed):
    P50: 12.058623ms
    P95: 17.645567ms
    P99: 208.142335ms

-- Consumer Group 7 --
  Reads Completed:   2994700
  Read Errors:       1
  Updates Completed: 2994700
  Claim Errors:      1
  Empty Claims:      99
  Avg Poll Size:     140.98
  Avg Throughput:    24955.83 reads/sec

  Read Latency (claim txn):
    P50: 7.008255ms
    P95: 7.806975ms
    P99: 8.364031ms

  End-to-End Latency (created_at → consumed):
    P50: 11.165695ms
    P95: 15.130623ms
    P99: 206.700543ms

-- Consumer Group 8 --
  Reads Completed:   2994700
  Read Errors:       1
  Updates Completed: 2994700
  Claim Errors:      0
  Empty Claims:      534
  Avg Poll Size:     144.90
  Avg Throughput:    24955.83 reads/sec

  Read Latency (claim txn):
    P50: 6.905855ms
    P95: 7.823359ms
    P99: 8.347647ms

  End-to-End Latency (created_at → consumed):
    P50: 11.411455ms
    P95: 17.367039ms
    P99: 205.783039ms

-- Consumer Group 9 --
  Reads Completed:   2994800
  Read Errors:       1
  Updates Completed: 2994800
  Claim Errors:      1
  Empty Claims:      1959
  Avg Poll Size:     152.89
  Avg Throughput:    24956.67 reads/sec

  Read Latency (claim txn):
    P50: 6.737919ms
    P95: 7.733247ms
    P99: 8.269823ms

  End-to-End Latency (created_at → consumed):
    P50: 11.902975ms
    P95: 17.825791ms
    P99: 207.093759ms

2025/10/02 15:05:56 pubsub benchmark complete
```