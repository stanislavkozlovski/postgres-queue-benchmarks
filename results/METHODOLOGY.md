# General Test Methodology

1. Run at ~50-60% CPU steady-state, depending on spikes. Get a stable run
2. Run for 2 minutes. Do 3 runs and average the results.
3. Run with a small, realistic machine (4 vCPU) purely default settings, then test with tuning
4. Run with the largest machine (e.g. 96 vCPU) to get a sense of the upper limits; with tuning.

Benchmarks are always imperfect (noisy neighbours, unrealistic workload, etc.). This is how I ran these.

# Queues

- Always fetch single records (massive contention and overhead per statement), as a queue would.
- At least once semantics

# Pub Sub

- Have high read fan-out (5x)
- Max out a single topic-partition, then try multiple (eg 50)
- Strict ordering, at least once semantincs