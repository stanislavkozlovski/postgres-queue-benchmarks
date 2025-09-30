Getting stuck at ~12-14k reads. Let's see where the bottleneck is:

# SELECT
```bash
benchmark=# EXPLAIN (ANALYZE, BUFFERS)
SELECT id, payload, created_at
FROM queue
ORDER BY id
FOR UPDATE SKIP LOCKED
LIMIT 1;
                                                                 QUERY PLAN                                                                 
--------------------------------------------------------------------------------------------------------------------------------------------
 Limit  (cost=0.30..0.37 rows=1 width=54) (actual time=0.355..0.355 rows=0 loops=1)
   Buffers: shared hit=221
   ->  LockRows  (cost=0.30..68320.83 rows=1085301 width=54) (actual time=0.354..0.354 rows=0 loops=1)
         Buffers: shared hit=221
         ->  Index Scan using idx_queue_id on queue  (cost=0.30..57467.82 rows=1085301 width=54) (actual time=0.187..0.295 rows=23 loops=1)
               Buffers: shared hit=198
 Planning Time: 0.044 ms
 Execution Time: 0.366 ms
(8 rows)
```
The select statement is pretty healthy.
ChatGPT is suggesting the WAL may be the bottleneck here, as each txn consists of a lot of operations
```

```