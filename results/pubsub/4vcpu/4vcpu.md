

```bash
ALTER TABLE public.topicpartition SET (
  autovacuum_vacuum_scale_factor = 0,
  autovacuum_analyze_scale_factor = 0.05
);
```