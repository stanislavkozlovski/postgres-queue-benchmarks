- server: c7i.xlarge
  - io2 40GB with 10,000 IOPS
  - Ubuntu 24.04
  - absolute default postgres settings
- client: c7i.xlarge
  - 50 writers
  - 50 readers
  - 1 KB payload
  - 120s duration

# Run 1 (unlimited writes)

```bash
./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres   --writers=50   --readers=50   --duration=120s   --payload=1024   --report=5s
```
Ran three times. See detailed results in [unlimited_writes_runs.md](./unlimited_writes_runs.md). It averages out to:

- ✍️ write: **12.2 MB/s**
- 📖️ read: **2.7 MB/s**
- write latency:
  - p99: 24 ms
  - p95: 10 ms
  - p50: 3.1 ms
- read latency:
  - p99: 73 ms
  - p95: 41 ms
  - p50: 15.5 ms
- bottleneck: 100% CPU

# Run 2 (limit writes to max reads we can get)

Since writes massively outpace reads, let's see if artificially throttling them to some rate can improve our reads.
Our goal is to figure out what's the max 1:1 ratio they can reach.  The methodology I follow is to just run with a few different max write rates, until I feel it is stable at a certain throughput.  I landed at 4.5 MB/s.

```bash
./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres   --writers=50   --readers=50   --duration=120s   --payload=1024   --report=5s --throttle_writes 4500
```
Ran three times. See detailed results in [limited_writes_runs.md](./limited_writes_runs.md). It averages out to:
