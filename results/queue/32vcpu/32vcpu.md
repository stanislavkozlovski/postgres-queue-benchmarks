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

Perhaps it's time to reconfigure postgres. It's been said numerous times before that it's default settings are very conservative:
> The Postgres defaults are very conservative. It will start even on a tiny machine like Raspberry Pi. Which is great! The flip side is it’s terrible for typical database servers which tend to have much more RAM/CPU. To get good performance on such larger machines, you need to adjust a couple parameters (shared_buffers, max_wal_size, …).


# Modified Settings
