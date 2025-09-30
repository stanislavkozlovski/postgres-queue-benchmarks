
![4vcpu_htop_screen.png](./4vcpu_htop_screen.png) for more details.


# Run 1:

```bash
ubuntu@ip-172-31-23-223:/tmp/postgres-queue-benchmarks$ ./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres   --writers=50   --readers=50   --duration=120s   --payload=1024   --report=5s
[12:37:40] W: 13444/s R: 4296/s QDepth: 45742 Err(W/R): 0/0
[12:37:45] W: 13994/s R: 4065/s QDepth: 95389 Err(W/R): 0/0
[12:37:50] W: 13213/s R: 3736/s QDepth: 142775 Err(W/R): 0/0
[12:37:55] W: 13225/s R: 3551/s QDepth: 191145 Err(W/R): 0/0
[12:38:00] W: 12752/s R: 3351/s QDepth: 238149 Err(W/R): 0/0
[12:38:05] W: 12596/s R: 3148/s QDepth: 285387 Err(W/R): 0/0
[12:38:10] W: 11860/s R: 3035/s QDepth: 329512 Err(W/R): 0/0
[12:38:15] W: 11886/s R: 2948/s QDepth: 374201 Err(W/R): 0/0
[12:38:20] W: 13049/s R: 2654/s QDepth: 426172 Err(W/R): 0/0
[12:38:25] W: 13058/s R: 3523/s QDepth: 473845 Err(W/R): 0/0
[12:38:30] W: 12349/s R: 3339/s QDepth: 518897 Err(W/R): 0/0
[12:38:35] W: 9232/s R: 1419/s QDepth: 557963 Err(W/R): 0/0
[12:38:40] W: 8687/s R: 834/s QDepth: 597229 Err(W/R): 0/0
[12:38:45] W: 7078/s R: 685/s QDepth: 629198 Err(W/R): 0/0
[12:38:50] W: 7116/s R: 584/s QDepth: 661856 Err(W/R): 0/0
[12:38:55] W: 6973/s R: 517/s QDepth: 694133 Err(W/R): 0/0
[12:39:00] W: 6224/s R: 579/s QDepth: 722359 Err(W/R): 0/0
[12:39:05] W: 11045/s R: 2527/s QDepth: 764948 Err(W/R): 0/0
[12:39:10] W: 12520/s R: 2944/s QDepth: 812828 Err(W/R): 0/0
[12:39:15] W: 12286/s R: 2848/s QDepth: 860016 Err(W/R): 0/0
[12:39:20] W: 11700/s R: 2748/s QDepth: 904775 Err(W/R): 0/0
[12:39:25] W: 11682/s R: 2616/s QDepth: 950104 Err(W/R): 0/0
[12:39:30] W: 11019/s R: 2606/s QDepth: 992169 Err(W/R): 0/0

=== Summary ===
Total Writes: 1341944
Total Reads: 305180
Total Updates: 305180
Write Errors: 0
Read Errors: 6
Avg Write Throughput: 11182.87 rows/sec
Avg Read Throughput: 2543.17 rows/sec

Write Latencies:
  P50: 3.201023ms
  P95: 11.149311ms
  P99: 26.296319ms

Read Latencies:
  P50: 14.966783ms
  P95: 53.870591ms
  P99: 100.335615ms

2025/09/30 12:39:35 benchmark complete
```
# Run 2
```bash
ubuntu@ip-172-31-23-223:/tmp/postgres-queue-benchmarks$ ./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres   --writers=50   --readers=50   --duration=120s   --payload=1024   --report=5s
[12:40:08] W: 14031/s R: 4298/s QDepth: 48665 Err(W/R): 0/0
[12:40:13] W: 14627/s R: 4018/s QDepth: 101708 Err(W/R): 0/0
[12:40:18] W: 13690/s R: 3708/s QDepth: 151616 Err(W/R): 0/0
[12:40:23] W: 13740/s R: 3506/s QDepth: 202787 Err(W/R): 0/0
[12:40:28] W: 13505/s R: 3312/s QDepth: 253751 Err(W/R): 0/0
[12:40:33] W: 13812/s R: 3127/s QDepth: 307179 Err(W/R): 0/0
[12:40:38] W: 12589/s R: 3050/s QDepth: 354883 Err(W/R): 0/0
[12:40:43] W: 8287/s R: 1135/s QDepth: 390640 Err(W/R): 0/0
[12:40:48] W: 7398/s R: 792/s QDepth: 423671 Err(W/R): 0/0
[12:40:53] W: 10793/s R: 2043/s QDepth: 467421 Err(W/R): 0/0
[12:40:58] W: 13309/s R: 3293/s QDepth: 517500 Err(W/R): 0/0
[12:41:03] W: 11118/s R: 2418/s QDepth: 561001 Err(W/R): 0/0
[12:41:08] W: 12627/s R: 3065/s QDepth: 608808 Err(W/R): 0/0
[12:41:13] W: 12835/s R: 2948/s QDepth: 658244 Err(W/R): 0/0
[12:41:18] W: 12642/s R: 2834/s QDepth: 707289 Err(W/R): 0/0
[12:41:23] W: 12450/s R: 2705/s QDepth: 756015 Err(W/R): 0/0
[12:41:28] W: 12656/s R: 2598/s QDepth: 806304 Err(W/R): 0/0
[12:41:33] W: 12377/s R: 2509/s QDepth: 855646 Err(W/R): 0/0
[12:41:38] W: 11618/s R: 2496/s QDepth: 901259 Err(W/R): 0/0
[12:41:43] W: 11773/s R: 2413/s QDepth: 948060 Err(W/R): 0/0
[12:41:48] W: 12081/s R: 2343/s QDepth: 996751 Err(W/R): 0/0
[12:41:53] W: 11378/s R: 2323/s QDepth: 1042026 Err(W/R): 0/0
[12:41:58] W: 11192/s R: 2248/s QDepth: 1086746 Err(W/R): 0/0

=== Summary ===
Total Writes: 1447305
Total Reads: 324410
Total Updates: 324410
Write Errors: 0
Read Errors: 25
Avg Write Throughput: 12060.88 rows/sec
Avg Read Throughput: 2703.42 rows/sec

Write Latencies:
  P50: 3.096575ms
  P95: 9.748479ms
  P99: 22.806527ms

Read Latencies:
  P50: 15.654911ms
  P95: 36.077567ms
  P99: 69.140479ms

2025/09/30 12:42:02 benchmark complete
```
# Run 3
```bash
ubuntu@ip-172-31-23-223:/tmp/postgres-queue-benchmarks$ ./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres   --writers=50   --readers=50   --duration=120s   --payload=1024   --report=5s
[12:42:18] W: 13912/s R: 4321/s QDepth: 47957 Err(W/R): 0/0
[12:42:23] W: 14022/s R: 3781/s QDepth: 99162 Err(W/R): 0/0
[12:42:28] W: 14164/s R: 3895/s QDepth: 150505 Err(W/R): 0/0
[12:42:33] W: 13553/s R: 3916/s QDepth: 198690 Err(W/R): 0/0
[12:42:38] W: 13781/s R: 3648/s QDepth: 249353 Err(W/R): 0/0
[12:42:43] W: 13473/s R: 3445/s QDepth: 299491 Err(W/R): 0/0
[12:42:48] W: 12906/s R: 3283/s QDepth: 347609 Err(W/R): 0/0
[12:42:53] W: 12571/s R: 3135/s QDepth: 394787 Err(W/R): 0/0
[12:42:58] W: 13082/s R: 2978/s QDepth: 445308 Err(W/R): 0/0
[12:43:03] W: 12469/s R: 2903/s QDepth: 493137 Err(W/R): 0/0
[12:43:08] W: 12188/s R: 2807/s QDepth: 540041 Err(W/R): 0/0
[12:43:13] W: 12639/s R: 2676/s QDepth: 589857 Err(W/R): 0/0
[12:43:18] W: 12403/s R: 2599/s QDepth: 638877 Err(W/R): 0/0
[12:43:23] W: 12007/s R: 2513/s QDepth: 686347 Err(W/R): 0/0
[12:43:28] W: 11873/s R: 2437/s QDepth: 733524 Err(W/R): 0/0
[12:43:33] W: 12048/s R: 2368/s QDepth: 781927 Err(W/R): 0/0
[12:43:38] W: 11920/s R: 2314/s QDepth: 829960 Err(W/R): 0/0
[12:43:43] W: 11154/s R: 2270/s QDepth: 874377 Err(W/R): 0/0
[12:43:48] W: 11266/s R: 2218/s QDepth: 919619 Err(W/R): 0/0
[12:43:53] W: 11056/s R: 2192/s QDepth: 963940 Err(W/R): 0/0
[12:43:58] W: 11021/s R: 2132/s QDepth: 1008385 Err(W/R): 0/0
[12:44:03] W: 10984/s R: 2091/s QDepth: 1052853 Err(W/R): 0/0
[12:44:08] W: 10547/s R: 2065/s QDepth: 1095266 Err(W/R): 0/0

=== Summary ===
Total Writes: 1471892
Total Reads: 338139
Total Updates: 338139
Write Errors: 0
Read Errors: 10
Avg Write Throughput: 12265.77 rows/sec
Avg Read Throughput: 2817.82 rows/sec

Write Latencies:
  P50: 3.059711ms
  P95: 9.437183ms
  P99: 22.724607ms

Read Latencies:
  P50: 15.728639ms
  P95: 31.899647ms
  P99: 48.955391ms

2025/09/30 12:44:12 benchmark complete
```