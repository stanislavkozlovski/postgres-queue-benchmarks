- server: c7i.8xlarge
    - io2 50GB with 12,000 IOPS
    - Ubuntu 24.04
    - absolute default postgres settings
    - 172.31.18.113
- client: c7i.xlarge
    - 50 writers
    - 50 readers
    - 1 KB payload
    - 120s duration

# Default Settings

- ✍️ write: **12.9 MB/s**
- 📖️ read: **12.9 MB/s**
- write latency:
  - p99: 1.7 ms
  - p95: 1.2 ms
  - p50: 0.8 ms
- read latency:
  - p99: 4.7 ms
  - p95: 3.6 ms
  - p50: 2.5 ms
- end-to-end latency:
  - p99: 1133 ms
  - p95: 633 ms
  - p50: 3.1 ms
- bottleneck: ? (CPU is at 50%, it must be something else)

Perhaps it's time to reconfigure postgres. It's well known online that its default settings are very conservative:
> The Postgres defaults are very conservative. It will start even on a tiny machine like Raspberry Pi. Which is great! The flip side is it’s terrible for typical database servers which tend to have much more RAM/CPU. To get good performance on such larger machines, you need to adjust a couple parameters (shared_buffers, max_wal_size, …).

We are running on a 32 vCPU machine with 64GB of RAM. Let's reconfigure:

# Modified Settings

I talked to ChatGPT a ton and it gave me the following configs:
- [vm dirty bytes kernel settings](./kernel_settings) - it believed WAL writes may have been a bottleneck, so suggested I slightly alter the pagecache dirty settings
- [postgres config](./modified_postgresql.conf) - it suggested I substantially increase some settings like `shared_buffers` and `max_worker_processes`.
- `queue` [table-specific vacuum settings](./table_vacuum_tuning.md) - it suggested more aggressive vacuuming for the `queue` table

This got me to ~16 MB/s, but I wasn't able to push further despite server CPU being low.
The next bottleneck was the **clients**! Postgres' average read latency was ~3ms, and I had 50 reader clients. Each client therefore can't sequentially execute more than 333 reads a second.
Collectively they can't pass 16650 req/s - ~16 MiB/s.
This comes from the simple fact that we're not batching. There is a lot of overhead (both in the client and server side) that each statement is only processing one message.
In the later tests, we will attempt batching. Systems like Kafka make heavy, heavy use of it.
At some point, the WAL fsync() calls will become the bottleneck. Per-commit durability, esp. when you have many thousands of commits a second, is expensive. I managed to get avg fsync down fron ~1.8ms to ~0.3ms by tuning the commit delay to 20 microseconds.

For now, let's just add more clients and see where this takes us.
Increasing the number of connections to 200 total (100 readers/100 writes) led its CPU pegged. Time for a bigger machine!
I re-deployed the client on a larger machine.

# Larger Client & Modified Settings

- server: c7i.8xlarge
  - io2 50GB with 12,000 IOPS
  - Ubuntu 24.04
  - modified settings
  - 172.31.18.113
- client: c7i.4xlarge
  - 100 writers
  - 100 readers
  - 1 KB payload
  - 120s duration

**_... TO BE CONTINUED ..._**