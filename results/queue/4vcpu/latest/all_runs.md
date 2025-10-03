# Run 1
```bash
ubuntu@ip-172-31-88-36:/tmp/postgres-queue-benchmarks$ ./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres   --writers=10   --readers=15   --duration=120s   --payload=1024   --report=5s   --throttle_writes=2900   --mode=queue
[11:41:40] W: 3472/s R: 3279/s QDepth: 965 Err(W/R): 0/0
[11:41:45] W: 2900/s R: 3090/s QDepth: 11 Err(W/R): 0/0
[11:41:50] W: 2901/s R: 2900/s QDepth: 14 Err(W/R): 0/0
[11:41:55] W: 2900/s R: 2901/s QDepth: 9 Err(W/R): 0/0
[11:42:00] W: 2900/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:42:05] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:42:10] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:42:15] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:42:20] W: 2900/s R: 2900/s QDepth: 11 Err(W/R): 0/0
[11:42:25] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:42:30] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:42:35] W: 2900/s R: 2900/s QDepth: 11 Err(W/R): 0/0
[11:42:40] W: 2900/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:42:45] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:42:50] W: 2899/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:42:55] W: 2901/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:43:00] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:43:05] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:43:10] W: 2900/s R: 2900/s QDepth: 14 Err(W/R): 0/0
[11:43:15] W: 2900/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:43:20] W: 2900/s R: 2900/s QDepth: 11 Err(W/R): 0/0
[11:43:25] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:43:30] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0

=== Summary ===
Total Writes: 350813
Total Reads: 350805
Total Updates: 350805
Write Errors: 0
Read Errors: 1
Avg Write Throughput: 2923.44 rows/sec
Avg Read Throughput: 2923.38 rows/sec

Write Latencies (INSERT only):
  P50: 1.880063ms
  P95: 2.381823ms
  P99: 2.539519ms

Read Latencies (txn: SELECT+DELETE+INSERT in queue; txn: UPDATE+SELECT range in pub-sub kafka semantics):
  P50: 4.268031ms
  P95: 5.177343ms
  P99: 5.500927ms

End-to-End Latencies (created_at → consumed):
  P50: 5.758975ms
  P95: 291.504127ms
  P99: 606.076927ms

2025/10/03 11:43:35 queue benchmark complete
```

# Run 2
```bash
 ./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres   --writers=10   --readers=15   --duration=120s   --payload=1024   --report=5s   --throttle_writes=2900   --mode=queue
[11:41:40] W: 3472/s R: 3279/s QDepth: 965 Err(W/R): 0/0
[11:41:45] W: 2900/s R: 3090/s QDepth: 11 Err(W/R): 0/0
[11:41:50] W: 2901/s R: 2900/s QDepth: 14 Err(W/R): 0/0
[11:41:55] W: 2900/s R: 2901/s QDepth: 9 Err(W/R): 0/0
[11:42:00] W: 2900/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:42:05] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:42:10] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:42:15] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:42:20] W: 2900/s R: 2900/s QDepth: 11 Err(W/R): 0/0
[11:42:25] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:42:30] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:42:35] W: 2900/s R: 2900/s QDepth: 11 Err(W/R): 0/0
[11:42:40] W: 2900/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:42:45] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:42:50] W: 2899/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:42:55] W: 2901/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:43:00] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:43:05] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:43:10] W: 2900/s R: 2900/s QDepth: 14 Err(W/R): 0/0
[11:43:15] W: 2900/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:43:20] W: 2900/s R: 2900/s QDepth: 11 Err(W/R): 0/0
[11:43:25] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:43:30] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0

=== Summary ===
Total Writes: 350813
Total Reads: 350805
Total Updates: 350805
Write Errors: 0
Read Errors: 1
Avg Write Throughput: 2923.44 rows/sec
Avg Read Throughput: 2923.38 rows/sec

Write Latencies (INSERT only):
  P50: 1.880063ms
  P95: 2.381823ms
  P99: 2.539519ms

Read Latencies (txn: SELECT+DELETE+INSERT in queue; txn: UPDATE+SELECT range in pub-sub kafka semantics):
  P50: 4.268031ms
  P95: 5.177343ms
  P99: 5.500927ms

End-to-End Latencies (created_at → consumed):
  P50: 5.758975ms
  P95: 291.504127ms
  P99: 606.076927ms

2025/10/03 11:43:35 queue benchmark complete
ubuntu@ip-172-31-88-36:/tmp/postgres-queue-benchmarks$ ./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres   --writers=10   --readers=15   --duration=120s   --payload=1024   --report=5s   --throttle_writes=2900   --mode=queue
[11:45:49] W: 3471/s R: 3313/s QDepth: 791 Err(W/R): 0/0
[11:45:54] W: 2901/s R: 3057/s QDepth: 12 Err(W/R): 0/0
[11:45:59] W: 2900/s R: 2900/s QDepth: 11 Err(W/R): 0/0
[11:46:04] W: 2900/s R: 2900/s QDepth: 8 Err(W/R): 0/0
[11:46:09] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:46:14] W: 2900/s R: 2901/s QDepth: 9 Err(W/R): 0/0
[11:46:19] W: 2900/s R: 2899/s QDepth: 12 Err(W/R): 0/0
[11:46:24] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:46:29] W: 2900/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:46:34] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:46:39] W: 2900/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:46:44] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:46:49] W: 2900/s R: 2800/s QDepth: 512 Err(W/R): 0/0
[11:46:54] W: 2900/s R: 3001/s QDepth: 9 Err(W/R): 0/0
[11:46:59] W: 2901/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:47:04] W: 2900/s R: 2900/s QDepth: 14 Err(W/R): 0/0
[11:47:09] W: 2900/s R: 2901/s QDepth: 11 Err(W/R): 0/0
[11:47:14] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:47:19] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:47:24] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:47:29] W: 2900/s R: 2899/s QDepth: 13 Err(W/R): 0/0
[11:47:34] W: 2900/s R: 2900/s QDepth: 11 Err(W/R): 0/0
[11:47:39] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0

=== Summary ===
Total Writes: 350752
Total Reads: 350743
Total Updates: 350743
Write Errors: 0
Read Errors: 3
Avg Write Throughput: 2922.93 rows/sec
Avg Read Throughput: 2922.86 rows/sec

Write Latencies (INSERT only):
  P50: 1.886207ms
  P95: 2.392063ms
  P99: 2.549759ms

Read Latencies (txn: SELECT+DELETE+INSERT in queue; txn: UPDATE+SELECT range in pub-sub kafka semantics):
  P50: 4.280319ms
  P95: 5.144575ms
  P99: 5.484543ms

End-to-End Latencies (created_at → consumed):
  P50: 5.779455ms
  P95: 252.182527ms
  P99: 574.619647ms

2025/10/03 11:47:44 queue benchmark complete
```
# Run 3
```bash
ubuntu@ip-172-31-88-36:/tmp/postgres-queue-benchmarks$ ./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres   --writers=10   --readers=15   --duration=120s   --payload=1024   --report=5s   --throttle_writes=2900   --mode=queue
 

[11:48:17] W: 3474/s R: 3367/s QDepth: 533 Err(W/R): 0/0
[11:48:22] W: 2900/s R: 3004/s QDepth: 11 Err(W/R): 0/0
[11:48:27] W: 2900/s R: 2900/s QDepth: 9 Err(W/R): 0/0
[11:48:32] W: 2900/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:48:37] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:48:42] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:48:47] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:48:52] W: 2900/s R: 2900/s QDepth: 11 Err(W/R): 0/0
[11:48:57] W: 2900/s R: 2900/s QDepth: 9 Err(W/R): 0/0
[11:49:02] W: 2900/s R: 2900/s QDepth: 9 Err(W/R): 0/0
[11:49:07] W: 2900/s R: 2900/s QDepth: 9 Err(W/R): 0/0
[11:49:12] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:49:17] W: 2900/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:49:22] W: 2900/s R: 2900/s QDepth: 10 Err(W/R): 0/0
[11:49:27] W: 2900/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:49:32] W: 2900/s R: 2901/s QDepth: 9 Err(W/R): 0/0
[11:49:37] W: 2900/s R: 2900/s QDepth: 9 Err(W/R): 0/0
[11:49:42] W: 2900/s R: 2900/s QDepth: 12 Err(W/R): 0/0
[11:49:47] W: 2900/s R: 2900/s QDepth: 9 Err(W/R): 0/0
[11:49:52] W: 2901/s R: 2898/s QDepth: 20 Err(W/R): 0/0
[11:49:57] W: 2900/s R: 2901/s QDepth: 12 Err(W/R): 0/0
[11:50:02] W: 2900/s R: 2900/s QDepth: 13 Err(W/R): 0/0
[11:50:07] W: 2900/s R: 2901/s QDepth: 11 Err(W/R): 0/0

=== Summary ===
Total Writes: 350765
Total Reads: 350756
Total Updates: 350756
Write Errors: 0
Read Errors: 3
Avg Write Throughput: 2923.04 rows/sec
Avg Read Throughput: 2922.97 rows/sec

Write Latencies (INSERT only):
  P50: 1.882111ms
  P95: 2.387967ms
  P99: 2.535423ms

Read Latencies (txn: SELECT+DELETE+INSERT in queue; txn: UPDATE+SELECT range in pub-sub kafka semantics):
  P50: 4.243455ms
  P95: 5.058559ms
  P99: 5.394431ms

End-to-End Latencies (created_at → consumed):
  P50: 5.681151ms
  P95: 100.532223ms
  P99: 538.443775ms

2025/10/03 11:50:12 queue benchmark complete
```