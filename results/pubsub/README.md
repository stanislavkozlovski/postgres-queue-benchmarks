```bash
export HOST="172.31.83.119"  # adjust to your server's private IP
```bash
./pg_queue_bench \
  --host=$HOST \
  --port=5432 \
  --db=benchmark \
  --user=postgres \
  --password=postgres \
  --duration=120s \
  --payload=1024 \
  --report=5s \
  --write-batch 100 \
  --read-batch 100 \
  --kafka-pub-sub-semantics \
  --writers=50 \
  --readers=2 \
  --consumer-groups 5 \
  --mode=pubsub
```
