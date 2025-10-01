I modified the benchmark to modify the vacuum settings based on a boolean flag. 
```SQL
ALTER TABLE queue SET (
    autovacuum_vacuum_scale_factor = 0.01,
    autovacuum_vacuum_insert_threshold = 1000,
    autovacuum_analyze_scale_factor = 0.05,
    fillfactor = 70
)
```
