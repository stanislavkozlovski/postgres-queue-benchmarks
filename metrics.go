package postgres_queue_benchmarks

import (
	"sync/atomic"

	"github.com/HdrHistogram/hdrhistogram-go"
)

type Metrics struct {
	WritesCompleted  atomic.Int64
	ReadsCompleted   atomic.Int64
	UpdatesCompleted atomic.Int64
	WriteErrors      atomic.Int64
	ReadErrors       atomic.Int64

	writerHists []*hdrhistogram.Histogram
	readerHists []*hdrhistogram.Histogram
}

func NewMetrics(writers, readers int) *Metrics {
	m := &Metrics{
		writerHists: make([]*hdrhistogram.Histogram, writers),
		readerHists: make([]*hdrhistogram.Histogram, readers),
	}
	for i := range m.writerHists {
		m.writerHists[i] = hdrhistogram.New(1, int64(10e9), 3) // 1ns–10s
	}
	for i := range m.readerHists {
		m.readerHists[i] = hdrhistogram.New(1, int64(10e9), 3)
	}
	return m
}
