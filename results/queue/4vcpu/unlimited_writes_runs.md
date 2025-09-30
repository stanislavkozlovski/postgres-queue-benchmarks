The main problem here is that writes outpace reads by a ton, so the system perpetually accumulates backlog.

# Run 1:

```bash
ubuntu@ip-172-31-23-223:/tmp/postgres-queue-benchmarks$ for i in {1..3}; do ./pg_queue_bench --host=$HOST --port=5432 --db=benchmark --user=postgres --password=postgres --writers=50 --readers=50 --duration=120s --payload=1024 --report=5s; [ $i -lt 3 ] && sleep 120; done
[13:50:26] W: 13288/s R: 4365/s QDepth: 44616 Err(W/R): 0/0
[13:50:31] W: 13939/s R: 4059/s QDepth: 94016 Err(W/R): 0/0
[13:50:36] W: 13498/s R: 3632/s QDepth: 143346 Err(W/R): 0/0
[13:50:41] W: 13072/s R: 3574/s QDepth: 190837 Err(W/R): 0/0
[13:50:46] W: 13090/s R: 3376/s QDepth: 239406 Err(W/R): 0/0
[13:50:51] W: 12757/s R: 3222/s QDepth: 287080 Err(W/R): 0/0
[13:50:56] W: 13105/s R: 3095/s QDepth: 337132 Err(W/R): 0/0
[13:51:01] W: 12631/s R: 2980/s QDepth: 385389 Err(W/R): 0/0
[13:51:06] W: 12829/s R: 2846/s QDepth: 435305 Err(W/R): 0/0
[13:51:11] W: 12613/s R: 2758/s QDepth: 484583 Err(W/R): 0/0
[13:51:16] W: 11747/s R: 2711/s QDepth: 529764 Err(W/R): 0/0
[13:51:21] W: 12059/s R: 2605/s QDepth: 577034 Err(W/R): 0/0
[13:51:26] W: 11882/s R: 2513/s QDepth: 623881 Err(W/R): 0/0
[13:51:31] W: 11362/s R: 2451/s QDepth: 668436 Err(W/R): 0/0
[13:51:36] W: 11640/s R: 2380/s QDepth: 714735 Err(W/R): 0/0
[13:51:41] W: 11359/s R: 2319/s QDepth: 759940 Err(W/R): 0/0
[13:51:46] W: 11271/s R: 2288/s QDepth: 804852 Err(W/R): 0/0
[13:51:51] W: 11318/s R: 2221/s QDepth: 850338 Err(W/R): 0/0
[13:51:56] W: 10607/s R: 2175/s QDepth: 892498 Err(W/R): 0/0
[13:52:01] W: 10552/s R: 2141/s QDepth: 934551 Err(W/R): 0/0
[13:52:06] W: 10257/s R: 2096/s QDepth: 975358 Err(W/R): 0/0
[13:52:11] W: 12652/s R: 2761/s QDepth: 1024813 Err(W/R): 0/0
[13:52:16] W: 12428/s R: 3038/s QDepth: 1071763 Err(W/R): 0/0

=== Summary ===
Total Writes: 1459712
Total Reads: 341822
Total Updates: 341822
Write Errors: 0
Read Errors: 26
Avg Write Throughput: 12164.27 rows/sec
Avg Read Throughput: 2848.52 rows/sec
2025/09/30 13:52:22 [merge] dropped 78583 values outside histogram range

Write Latencies (INSERT only):
  P50: 3.129343ms
  P95: 9.420799ms
  P99: 23.347199ms

Read Latencies (txn: SELECT+DELETE+INSERT):
  P50: 15.859711ms
  P95: 31.522815ms
  P99: 47.218687ms

End-to-End Latencies (created_at → consumed):
  P50: 26.977763327s
  P95: 1m3.954747391s
  P99: 1m7.779952639s

2025/09/30 13:52:22 benchmark complete
```
# Run 2
```bash
[13:54:27] W: 14918/s R: 4268/s QDepth: 53247 Err(W/R): 0/0
[13:54:32] W: 14355/s R: 4061/s QDepth: 104717 Err(W/R): 0/0
[13:54:37] W: 14352/s R: 3592/s QDepth: 158514 Err(W/R): 0/0
[13:54:42] W: 13690/s R: 3539/s QDepth: 209271 Err(W/R): 0/0
[13:54:47] W: 12957/s R: 3406/s QDepth: 257026 Err(W/R): 0/0
[13:54:52] W: 13189/s R: 3226/s QDepth: 306843 Err(W/R): 0/0
[13:54:57] W: 12602/s R: 3098/s QDepth: 354363 Err(W/R): 0/0
[13:55:02] W: 13005/s R: 2966/s QDepth: 404560 Err(W/R): 0/0
[13:55:07] W: 12524/s R: 2887/s QDepth: 452744 Err(W/R): 0/0
[13:55:12] W: 12483/s R: 2775/s QDepth: 501286 Err(W/R): 0/0
[13:55:17] W: 12363/s R: 2706/s QDepth: 549572 Err(W/R): 0/0
[13:55:22] W: 12483/s R: 2580/s QDepth: 599089 Err(W/R): 0/0
[13:55:27] W: 11657/s R: 2521/s QDepth: 644769 Err(W/R): 0/0
[13:55:32] W: 11725/s R: 2443/s QDepth: 691177 Err(W/R): 0/0
[13:55:37] W: 11876/s R: 2386/s QDepth: 738629 Err(W/R): 0/0
[13:55:42] W: 12008/s R: 2298/s QDepth: 787176 Err(W/R): 0/0
[13:55:47] W: 11483/s R: 2273/s QDepth: 833230 Err(W/R): 0/0
[13:55:52] W: 11537/s R: 2231/s QDepth: 879759 Err(W/R): 0/0
[13:55:57] W: 11420/s R: 2176/s QDepth: 925978 Err(W/R): 0/0
[13:56:02] W: 11405/s R: 2141/s QDepth: 972300 Err(W/R): 0/0
[13:56:07] W: 11106/s R: 2094/s QDepth: 1017359 Err(W/R): 0/0
[13:56:12] W: 13411/s R: 2683/s QDepth: 1070997 Err(W/R): 0/0
[13:56:17] W: 13316/s R: 3025/s QDepth: 1122451 Err(W/R): 0/0

=== Summary ===
Total Writes: 1510928
Total Reads: 340934
Total Updates: 340934
Write Errors: 0
Read Errors: 23
Avg Write Throughput: 12591.07 rows/sec
Avg Read Throughput: 2841.12 rows/sec
2025/09/30 13:56:22 [merge] dropped 82228 values outside histogram range

Write Latencies (INSERT only):
  P50: 2.985983ms
  P95: 9.199615ms
  P99: 23.085055ms

Read Latencies (txn: SELECT+DELETE+INSERT):
  P50: 15.908863ms
  P95: 31.752191ms
  P99: 48.201727ms

End-to-End Latencies (created_at → consumed):
  P50: 27.430748159s
  P95: 1m3.921192959s
  P99: 1m7.746398207s

2025/09/30 13:56:22 benchmark complete
```
# Run 3
```bash
[13:58:27] W: 13167/s R: 4422/s QDepth: 43727 Err(W/R): 0/0
[13:58:32] W: 14246/s R: 4094/s QDepth: 94487 Err(W/R): 0/0
[13:58:37] W: 14293/s R: 3516/s QDepth: 148371 Err(W/R): 0/0
[13:58:42] W: 13625/s R: 3606/s QDepth: 198465 Err(W/R): 0/0
[13:58:47] W: 13119/s R: 3434/s QDepth: 246894 Err(W/R): 0/0
[13:58:52] W: 13279/s R: 3252/s QDepth: 297027 Err(W/R): 0/0
[13:58:57] W: 12833/s R: 3092/s QDepth: 345732 Err(W/R): 0/0
[13:59:02] W: 12727/s R: 2980/s QDepth: 394466 Err(W/R): 0/0
[13:59:07] W: 12490/s R: 2870/s QDepth: 442566 Err(W/R): 0/0
[13:59:12] W: 12549/s R: 2772/s QDepth: 491453 Err(W/R): 0/0
[13:59:17] W: 12479/s R: 2679/s QDepth: 540453 Err(W/R): 0/0
[13:59:22] W: 12337/s R: 2578/s QDepth: 589250 Err(W/R): 0/0
[13:59:27] W: 12152/s R: 2502/s QDepth: 637500 Err(W/R): 0/0
[13:59:32] W: 11834/s R: 2439/s QDepth: 684473 Err(W/R): 0/0
[13:59:37] W: 11734/s R: 2366/s QDepth: 731312 Err(W/R): 0/0
[13:59:42] W: 11952/s R: 2289/s QDepth: 779626 Err(W/R): 0/0
[13:59:47] W: 11929/s R: 2243/s QDepth: 828054 Err(W/R): 0/0
[13:59:52] W: 11637/s R: 2220/s QDepth: 875138 Err(W/R): 0/0
[13:59:57] W: 11156/s R: 2183/s QDepth: 920005 Err(W/R): 0/0
[14:00:02] W: 11284/s R: 2134/s QDepth: 965756 Err(W/R): 0/0
[14:00:07] W: 11146/s R: 2082/s QDepth: 1011077 Err(W/R): 0/0
[14:00:12] W: 13002/s R: 2842/s QDepth: 1061876 Err(W/R): 0/0
[14:00:17] W: 13259/s R: 3009/s QDepth: 1113127 Err(W/R): 0/0

=== Summary ===
Total Writes: 1504545
Total Reads: 342219
Total Updates: 342219
Write Errors: 0
Read Errors: 21
Avg Write Throughput: 12537.88 rows/sec
Avg Read Throughput: 2851.82 rows/sec
2025/09/30 14:00:22 [merge] dropped 80771 values outside histogram range

Write Latencies (INSERT only):
  P50: 3.008511ms
  P95: 9.297919ms
  P99: 23.117823ms

Read Latencies (txn: SELECT+DELETE+INSERT):
  P50: 15.818751ms
  P95: 31.621119ms
  P99: 47.710207ms

End-to-End Latencies (created_at → consumed):
  P50: 26.625441791s
  P95: 1m3.854084095s
  P99: 1m7.712843775s

2025/09/30 14:00:22 benchmark complete
```