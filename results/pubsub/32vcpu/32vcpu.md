
- us-east-1c
- c7i.4xlarge
- goal ~60CPU
  - c6i.16xlarge


37.5 gibps (gigabits per second) is equal to 4.6875 gib per second (GiB/s) or ~4,687,500,000 MiB/s (mebibytes per second)
c7i.24xlarge

```bash
./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres \
 --report=5s    --write-batch=100   --read-batch=200 \
    --writers=50   --readers=20   --consumer-groups=5    \
    --mode=pubsub --partitions=20 --duration=120s --throttle_writes=50000
```