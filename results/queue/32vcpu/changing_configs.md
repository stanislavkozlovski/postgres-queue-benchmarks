Create a bash file
```bash
#!/bin/bash
CONF="/etc/postgresql/17/main/postgresql.conf"

# list of params to comment out
PARAMS=(
  max_connections
  shared_buffers
  effective_cache_size
  work_mem
  maintenance_work_mem
  huge_pages
  wal_level
  wal_compression
  wal_buffers
  min_wal_size
  max_wal_size
  checkpoint_timeout
  checkpoint_completion_target
  synchronous_commit
  default_statistics_target
  random_page_cost
  effective_io_concurrency
  max_worker_processes
  max_parallel_workers
  max_parallel_workers_per_gather
  max_parallel_maintenance_workers
  autovacuum_max_workers
  autovacuum_naptime
  autovacuum_vacuum_cost_limit
  autovacuum_vacuum_cost_delay
)

for p in "${PARAMS[@]}"; do
  sudo sed -i "s/^\s*\(${p}\s*=\s*.*\)$/# \1/" "$CONF"
done

echo "✅ All matching parameters commented out in $CONF"

```

Run it.
```bash
chmod +x comment_out.sh
./comment_out.sh 
```

Paste the [modified postgresql config](./modified_postgresql.conf) into `conf.d`:
```bash
   77  vi /etc/postgresql/17/main/conf.d/99-custom.conf
```

Restart PG
```bash
sudo systemctl reload postgresql@17-main
```



Check if settings are present
```bash
# psql into machine
SHOW shared_buffers;
```

Run the test with the new tune table vacuum flag
```
 ./pg_queue_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres  
 --password=postgres   --writers=50   --readers=50   --duration=120s   --payload=1024   --report=5s --throttle_writes 15000 \
 --tune-table-vacuum # <--NEW
```
