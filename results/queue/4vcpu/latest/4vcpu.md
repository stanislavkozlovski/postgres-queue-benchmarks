# 4 vCPU Results

- ✍️ write: **2.85 MB/s** (**2923 msg/s**)
- 📖️ read: **2.85 MB/s** (**2923 msg/s**)
- CPU: ~60%
- write latency:
  - p99: **2.54 ms**
  - p95: **2.39 ms**
  - p50: **1.88 ms**
- read latency:
  - p99: **5.46 ms**
  - p95: **5.13 ms**
  - p50: **4.26 ms**
- ❌ end-to-end latency:
  - p99: **0.573 s**
  - p95: **0.215 s**
  - p50: **0.006 s**

- server: c7i.xlarge (**the cheap version**)
    - gp3 25GB 8000 IOPS
    - Ubuntu 24.04
    - absolute default postgres settings
    - goal: ~60% CPU
    - $0.118/h (1yr reserved) - $1033.68 a year
    - assume ~12h message retention (then to s3 or whatever); - deploy a 242GB disk (19.36/m at $0.08/GB-month)
      - 6000 * 0.005/provisioned IOPS-month = 30/m = $592.32/yr
      - in reality, can offload to s3 by the hour
    - us-east-1a
- client: c7i.xlarge
    - 10 writers
    - 15 readers
    - 1 KB payload
    - 120s duration
    - us-east-1a

Notes:
- CPU bottlenecks due to number of readers/writers. With 50/50 read/write connections, it pegged to 100% CPU. At 100% CPU we could get ~4 MB/s out of the machine. And queue depth didn't seem to increase. But it was pegged at 100%.
- At this 4 vCPU size, we need less clients. 10 writes and 17 readers is a sweet spot. CPU at 60-68%.
- 3000k IOPS on gp3 was low too, had to increase it to 6000

# Run (limiting writes to max reads we can get)

Since writes massively outpace reads, we throttle writes to find the max 1:1 ratio the server can reach.
The bottleneck is the number of **reader clients**. If a client's p50 read latency is 5.4ms, then it can only read 185 messages a second.
The total read throughput is therefore bottlenecked by how many readers we can run in parallel without blowing up CPU.
As we said, we keep it at 60% CPU.
We limit ourselves to 15 readers at 2900 msg/s total.
```bash
 ./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres  \
 --writers=10   --readers=15   --duration=120s   --payload=1024   --report=5s \
 --throttle_writes=2900   --mode=queue
```

