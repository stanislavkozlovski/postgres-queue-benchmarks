# Run 1
```bash
ubuntu@ip-172-31-23-223:/tmp/postgres-queue-benchmarks$ ./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres   --writers=50   --readers=50   --duration=120s   --payload=1024   --report=5s --throttle_writes 4500
[12:59:02] W: 5390/s R: 5381/s QDepth: 45 Err(W/R): 0/0
[12:59:07] W: 4500/s R: 4501/s QDepth: 38 Err(W/R): 0/0
[12:59:12] W: 4500/s R: 4500/s QDepth: 35 Err(W/R): 0/0
[12:59:17] W: 4500/s R: 4498/s QDepth: 48 Err(W/R): 0/0
[12:59:22] W: 4501/s R: 4504/s QDepth: 31 Err(W/R): 0/0
[12:59:27] W: 4500/s R: 4499/s QDepth: 35 Err(W/R): 0/0
[12:59:32] W: 4500/s R: 4499/s QDepth: 39 Err(W/R): 0/0
[12:59:37] W: 4500/s R: 4501/s QDepth: 33 Err(W/R): 0/0
[12:59:42] W: 4500/s R: 4498/s QDepth: 42 Err(W/R): 0/0
[12:59:47] W: 4500/s R: 4501/s QDepth: 35 Err(W/R): 0/0
[12:59:52] W: 4499/s R: 4501/s QDepth: 28 Err(W/R): 0/0
[12:59:57] W: 4500/s R: 4499/s QDepth: 34 Err(W/R): 0/0
[13:00:02] W: 4501/s R: 4501/s QDepth: 33 Err(W/R): 0/0
[13:00:07] W: 4500/s R: 4499/s QDepth: 35 Err(W/R): 0/0
[13:00:12] W: 4499/s R: 4500/s QDepth: 29 Err(W/R): 0/0
[13:00:17] W: 4501/s R: 4500/s QDepth: 33 Err(W/R): 0/0
[13:00:22] W: 4500/s R: 4499/s QDepth: 34 Err(W/R): 0/0
[13:00:27] W: 4500/s R: 4495/s QDepth: 59 Err(W/R): 0/0
[13:00:32] W: 4499/s R: 4505/s QDepth: 29 Err(W/R): 0/0
[13:00:37] W: 4501/s R: 4501/s QDepth: 28 Err(W/R): 0/0
[13:00:42] W: 4500/s R: 3313/s QDepth: 5964 Err(W/R): 0/0
[13:00:47] W: 4499/s R: 5423/s QDepth: 1347 Err(W/R): 0/0
[13:00:52] W: 4501/s R: 4754/s QDepth: 83 Err(W/R): 0/0

=== Summary ===
Total Writes: 544302
Total Reads: 543949
Total Updates: 543949
Write Errors: 0
Read Errors: 11
Avg Write Throughput: 4535.85 rows/sec
Avg Read Throughput: 4532.91 rows/sec

Write Latencies:
  P50: 2.069503ms
  P95: 4.134911ms
  P99: 6.885375ms

Read Latencies:
  P50: 8.396799ms
  P95: 12.697599ms
  P99: 19.120127ms

2025/09/30 13:00:57 benchmark complete
```

# Run 2
```bash
```