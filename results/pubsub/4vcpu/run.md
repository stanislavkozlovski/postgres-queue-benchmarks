- default settings
- 4vcpu 8 gib ram
- 10k iops gp3

Less writers! 10 writers for example

```bash
 ./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres --report=5s  \
  --write-batch 100   --read-batch 200   --kafka-pub-sub-semantics   --writers=15   --readers=2   --consumer-groups 10   \ 
  --mode=pubsub --throttle_writes=25000 --duration=120s
```

Then upscale and go again

# Run 1

```bash

=== PubSub Summary ===
Delivery Semantics: at-least-once (strict ordering)
Duration: 2m0s (120.0s)

-- Consumer Group 0 --
  Reads Completed:   3019100
  Read Errors:       1
  Updates Completed: 3019100
  Claim Errors:      1
  Empty Claims:      297
  Avg Poll Size:     132.73
  Avg Throughput:    25159.17 reads/sec

  Read Latency (claim txn):
    P50: 6.549503ms
    P95: 7.307263ms
    P99: 7.835647ms

  End-to-End Latency (created_at → consumed):
    P50: 10.756095ms
    P95: 14.082047ms
    P99: 94.961663ms

-- Consumer Group 1 --
  Reads Completed:   3019100
  Read Errors:       1
  Updates Completed: 3019100
  Claim Errors:      1
  Empty Claims:      910
  Avg Poll Size:     148.22
  Avg Throughput:    25159.17 reads/sec

  Read Latency (claim txn):
    P50: 6.914047ms
    P95: 7.819263ms
    P99: 8.294399ms

  End-to-End Latency (created_at → consumed):
    P50: 11.501567ms
    P95: 16.662527ms
    P99: 206.176255ms

-- Consumer Group 2 --
  Reads Completed:   3019200
  Read Errors:       0
  Updates Completed: 3019200
  Claim Errors:      1
  Empty Claims:      1437
  Avg Poll Size:     127.13
  Avg Throughput:    25160.00 reads/sec

  Read Latency (claim txn):
    P50: 5.967871ms
    P95: 6.705151ms
    P99: 7.184383ms

  End-to-End Latency (created_at → consumed):
    P50: 10.321919ms
    P95: 14.385151ms
    P99: 43.057151ms

-- Consumer Group 3 --
  Reads Completed:   3019300
  Read Errors:       0
  Updates Completed: 3019300
  Claim Errors:      1
  Empty Claims:      2624
  Avg Poll Size:     146.02
  Avg Throughput:    25160.83 reads/sec

  Read Latency (claim txn):
    P50: 6.328319ms
    P95: 7.229439ms
    P99: 7.745535ms

  End-to-End Latency (created_at → consumed):
    P50: 11.370495ms
    P95: 16.326655ms
    P99: 67.764223ms

-- Consumer Group 4 --
  Reads Completed:   3019100
  Read Errors:       1
  Updates Completed: 3019100
  Claim Errors:      1
  Empty Claims:      1183
  Avg Poll Size:     150.11
  Avg Throughput:    25159.17 reads/sec

  Read Latency (claim txn):
    P50: 6.885375ms
    P95: 7.831551ms
    P99: 8.323071ms

  End-to-End Latency (created_at → consumed):
    P50: 11.583487ms
    P95: 16.957439ms
    P99: 200.671231ms

-- Consumer Group 5 --
  Reads Completed:   3019200
  Read Errors:       0
  Updates Completed: 3019200
  Claim Errors:      1
  Empty Claims:      2425
  Avg Poll Size:     153.07
  Avg Throughput:    25160.00 reads/sec

  Read Latency (claim txn):
    P50: 6.643711ms
    P95: 7.622655ms
    P99: 8.138751ms

  End-to-End Latency (created_at → consumed):
    P50: 11.780095ms
    P95: 16.891903ms
    P99: 158.728191ms

-- Consumer Group 6 --
  Reads Completed:   3019200
  Read Errors:       0
  Updates Completed: 3019200
  Claim Errors:      1
  Empty Claims:      282
  Avg Poll Size:     133.29
  Avg Throughput:    25160.00 reads/sec

  Read Latency (claim txn):
    P50: 6.594559ms
    P95: 7.331839ms
    P99: 7.856127ms

  End-to-End Latency (created_at → consumed):
    P50: 10.698751ms
    P95: 14.024703ms
    P99: 97.779711ms

-- Consumer Group 7 --
  Reads Completed:   3019300
  Read Errors:       0
  Updates Completed: 3019300
  Claim Errors:      1
  Empty Claims:      2133
  Avg Poll Size:     152.62
  Avg Throughput:    25160.83 reads/sec

  Read Latency (claim txn):
    P50: 6.696959ms
    P95: 7.663615ms
    P99: 8.130559ms

  End-to-End Latency (created_at → consumed):
    P50: 11.763711ms
    P95: 17.022975ms
    P99: 162.267135ms

-- Consumer Group 8 --
  Reads Completed:   3019100
  Read Errors:       1
  Updates Completed: 3019100
  Claim Errors:      1
  Empty Claims:      685
  Avg Poll Size:     143.18
  Avg Throughput:    25159.17 reads/sec

  Read Latency (claim txn):
    P50: 6.803455ms
    P95: 7.667711ms
    P99: 8.175615ms

  End-to-End Latency (created_at → consumed):
    P50: 11.255807ms
    P95: 15.876095ms
    P99: 174.194687ms

-- Consumer Group 9 --
  Reads Completed:   3019200
  Read Errors:       0
  Updates Completed: 3019200
  Claim Errors:      1
  Empty Claims:      260
  Avg Poll Size:     133.00
  Avg Throughput:    25160.00 reads/sec

  Read Latency (claim txn):
    P50: 6.582271ms
    P95: 7.315455ms
    P99: 7.806975ms

  End-to-End Latency (created_at → consumed):
    P50: 10.739711ms
    P95: 14.008319ms
    P99: 107.413503ms

=== Aggregate Across All Groups ===
  Total Reads Completed: 30191800
  Total Read Errors:     4
  Total Updates:         30191800
  Total Claim Errors:    10
  Total Empty Claims:    12236
  Avg Poll Size:         141.35
  Avg Throughput:        251598.33 reads/sec

  Read Latency (claim txn):
    P50: 6.569983ms
    P95: 7.565311ms
    P99: 8.073215ms

  End-to-End Latency (created_at → consumed):
    P50: 11.149311ms
    P95: 15.933439ms
    P99: 130.285567ms

2025/10/02 16:23:21 pubsub benchmark complete
```