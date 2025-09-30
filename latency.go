package postgres_queue_benchmarks

import (
	"math"
	"sort"
	"sync"
	"time"
)

type latencyBag struct {
	s  []time.Duration
	mu sync.Mutex
}

func (l *latencyBag) add(d time.Duration) {
	l.mu.Lock()
	l.s = append(l.s, d)
	l.mu.Unlock()
}

func (l *latencyBag) snapshot() []time.Duration {
	l.mu.Lock()
	cp := append([]time.Duration(nil), l.s...)
	l.mu.Unlock()
	return cp
}

func percentile(lat []time.Duration, p float64) time.Duration {
	if len(lat) == 0 {
		return 0
	}
	cp := append([]time.Duration(nil), lat...)
	sort.Slice(cp, func(i, j int) bool { return cp[i] < cp[j] })
	idx := int(math.Ceil((p/100.0)*float64(len(cp)))) - 1
	if idx < 0 {
		idx = 0
	}
	if idx >= len(cp) {
		idx = len(cp) - 1
	}
	return cp[idx]
}
