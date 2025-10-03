
```bash
./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres  \
 --writers=100   --readers=100   --duration=120s   --payload=1024   --report=5s \
 --throttle_writes=20000   --mode=queue
```
