```bash
./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark
   --user=postgres   --password=postgres   --writers=40   --readers=40   --duration=120s   --payload=1024   --report=5s --throttle_writes 
13000
```

![screen](./default_settings50_cpu_at_13k_writes.png)

# Run 1
```bash
[15:45:16] W: 15599/s R: 15248/s QDepth: 1759 Err(W/R): 0/0
[15:45:21] W: 13007/s R: 13352/s QDepth: 34 Err(W/R): 0/0
[15:45:26] W: 13004/s R: 13005/s QDepth: 25 Err(W/R): 0/0
[15:45:31] W: 13002/s R: 12999/s QDepth: 41 Err(W/R): 0/0
[15:45:36] W: 13002/s R: 13004/s QDepth: 31 Err(W/R): 0/0
[15:45:41] W: 13003/s R: 13003/s QDepth: 31 Err(W/R): 0/0
[15:45:46] W: 13005/s R: 13005/s QDepth: 34 Err(W/R): 0/0
[15:45:51] W: 13002/s R: 13003/s QDepth: 30 Err(W/R): 0/0
[15:45:56] W: 13002/s R: 13002/s QDepth: 29 Err(W/R): 0/0
[15:46:01] W: 13001/s R: 13000/s QDepth: 36 Err(W/R): 0/0
[15:46:06] W: 13003/s R: 13005/s QDepth: 23 Err(W/R): 0/0
[15:46:11] W: 13002/s R: 13002/s QDepth: 26 Err(W/R): 0/0
[15:46:16] W: 13003/s R: 13002/s QDepth: 32 Err(W/R): 0/0
[15:46:21] W: 13007/s R: 9859/s QDepth: 15773 Err(W/R): 0/0
[15:46:26] W: 13002/s R: 15244/s QDepth: 4564 Err(W/R): 0/0
[15:46:31] W: 13004/s R: 13912/s QDepth: 27 Err(W/R): 0/0
[15:46:36] W: 13004/s R: 13003/s QDepth: 32 Err(W/R): 0/0
[15:46:41] W: 13003/s R: 13003/s QDepth: 32 Err(W/R): 0/0
[15:46:46] W: 13004/s R: 13005/s QDepth: 27 Err(W/R): 0/0
[15:46:51] W: 13002/s R: 13000/s QDepth: 37 Err(W/R): 0/0
[15:46:56] W: 13006/s R: 13007/s QDepth: 32 Err(W/R): 0/0
[15:47:01] W: 13001/s R: 13003/s QDepth: 22 Err(W/R): 0/0
[15:47:06] W: 13003/s R: 13001/s QDepth: 31 Err(W/R): 0/0

=== Summary ===
Total Writes: 1559536
Total Reads: 1559514
Total Updates: 1559514
Write Errors: 0
Read Errors: 12
Avg Write Throughput: 12996.13 rows/sec
Avg Read Throughput: 12995.95 rows/sec

Write Latencies (INSERT only):
  P50: 781.823µs
  P95: 1.230847ms
  P99: 1.702911ms

Read Latencies (txn: SELECT+DELETE+INSERT):
  P50: 2.494463ms
  P95: 3.592191ms
  P99: 4.673535ms

End-to-End Latencies (created_at → consumed):
  P50: 3.158015ms
  P95: 644.874239ms
  P99: 1.126170623s

2025/09/30 15:47:10 benchmark complete
```
# Run 2
```bash

[15:49:15] W: 15582/s R: 15334/s QDepth: 1241 Err(W/R): 0/0
[15:49:20] W: 13004/s R: 13246/s QDepth: 29 Err(W/R): 0/0
[15:49:25] W: 13004/s R: 13003/s QDepth: 32 Err(W/R): 0/0
[15:49:30] W: 13009/s R: 13007/s QDepth: 42 Err(W/R): 0/0
[15:49:35] W: 13005/s R: 13007/s QDepth: 29 Err(W/R): 0/0
[15:49:40] W: 13003/s R: 13002/s QDepth: 32 Err(W/R): 0/0
[15:49:45] W: 13001/s R: 13002/s QDepth: 29 Err(W/R): 0/0
[15:49:50] W: 13003/s R: 13002/s QDepth: 33 Err(W/R): 0/0
[15:49:55] W: 13003/s R: 13003/s QDepth: 34 Err(W/R): 0/0
[15:50:00] W: 13002/s R: 13004/s QDepth: 23 Err(W/R): 0/0
[15:50:05] W: 13003/s R: 13003/s QDepth: 22 Err(W/R): 0/0
[15:50:10] W: 13007/s R: 13005/s QDepth: 29 Err(W/R): 0/0
[15:50:15] W: 13001/s R: 13002/s QDepth: 27 Err(W/R): 0/0
[15:50:20] W: 13003/s R: 11048/s QDepth: 9806 Err(W/R): 0/0
[15:50:25] W: 13004/s R: 13597/s QDepth: 6839 Err(W/R): 0/0
[15:50:30] W: 13002/s R: 14364/s QDepth: 28 Err(W/R): 0/0
[15:50:35] W: 13001/s R: 13003/s QDepth: 21 Err(W/R): 0/0
[15:50:40] W: 13004/s R: 13001/s QDepth: 33 Err(W/R): 0/0
[15:50:45] W: 13001/s R: 13002/s QDepth: 27 Err(W/R): 0/0
[15:50:50] W: 13002/s R: 13002/s QDepth: 27 Err(W/R): 0/0
[15:50:55] W: 13004/s R: 13003/s QDepth: 31 Err(W/R): 0/0
[15:51:00] W: 13004/s R: 13001/s QDepth: 48 Err(W/R): 0/0
[15:51:05] W: 13009/s R: 13013/s QDepth: 25 Err(W/R): 0/0

=== Summary ===
Total Writes: 1572274
Total Reads: 1572254
Total Updates: 1572254
Write Errors: 0
Read Errors: 9
Avg Write Throughput: 13102.28 rows/sec
Avg Read Throughput: 13102.12 rows/sec

Write Latencies (INSERT only):
  P50: 778.751µs
  P95: 1.233919ms
  P99: 1.712127ms

Read Latencies (txn: SELECT+DELETE+INSERT):
  P50: 2.476031ms
  P95: 3.588095ms
  P99: 4.665343ms

End-to-End Latencies (created_at → consumed):
  P50: 3.123199ms
  P95: 634.388479ms
  P99: 1.135607807s

2025/09/30 15:51:10 benchmark complete
```
# Run 3
```bash
[15:53:15] W: 15576/s R: 15481/s QDepth: 475 Err(W/R): 0/0
[15:53:20] W: 13003/s R: 13092/s QDepth: 32 Err(W/R): 0/0
[15:53:25] W: 13004/s R: 13004/s QDepth: 30 Err(W/R): 0/0
[15:53:30] W: 13005/s R: 13006/s QDepth: 23 Err(W/R): 0/0
[15:53:35] W: 13002/s R: 12999/s QDepth: 41 Err(W/R): 0/0
[15:53:40] W: 13001/s R: 13003/s QDepth: 32 Err(W/R): 0/0
[15:53:45] W: 13003/s R: 13004/s QDepth: 31 Err(W/R): 0/0
[15:53:50] W: 13002/s R: 13001/s QDepth: 35 Err(W/R): 0/0
[15:53:55] W: 13004/s R: 13006/s QDepth: 25 Err(W/R): 0/0
[15:54:00] W: 13004/s R: 12997/s QDepth: 62 Err(W/R): 0/0
[15:54:05] W: 13004/s R: 13009/s QDepth: 34 Err(W/R): 0/0
[15:54:10] W: 13002/s R: 13003/s QDepth: 29 Err(W/R): 0/0
[15:54:15] W: 13004/s R: 13003/s QDepth: 31 Err(W/R): 0/0
[15:54:20] W: 13005/s R: 9893/s QDepth: 15594 Err(W/R): 0/0
[15:54:25] W: 13004/s R: 15118/s QDepth: 5023 Err(W/R): 0/0
[15:54:30] W: 13003/s R: 14002/s QDepth: 29 Err(W/R): 0/0
[15:54:35] W: 13004/s R: 13004/s QDepth: 28 Err(W/R): 0/0
[15:54:40] W: 13008/s R: 13008/s QDepth: 28 Err(W/R): 0/0
[15:54:45] W: 13004/s R: 13004/s QDepth: 31 Err(W/R): 0/0
[15:54:50] W: 13004/s R: 13005/s QDepth: 27 Err(W/R): 0/0
[15:54:55] W: 13004/s R: 13003/s QDepth: 33 Err(W/R): 0/0
[15:55:00] W: 12999/s R: 12852/s QDepth: 766 Err(W/R): 0/0
[15:55:05] W: 13006/s R: 13152/s QDepth: 35 Err(W/R): 0/0

=== Summary ===
Total Writes: 1572276
Total Reads: 1572250
Total Updates: 1572250
Write Errors: 0
Read Errors: 16
Avg Write Throughput: 13102.30 rows/sec
Avg Read Throughput: 13102.08 rows/sec

Write Latencies (INSERT only):
  P50: 763.903µs
  P95: 1.213439ms
  P99: 1.693695ms

Read Latencies (txn: SELECT+DELETE+INSERT):
  P50: 2.459647ms
  P95: 3.561471ms
  P99: 4.653055ms

End-to-End Latencies (created_at → consumed):
  P50: 3.080191ms
  P95: 618.135551ms
  P99: 1.137704959s
```