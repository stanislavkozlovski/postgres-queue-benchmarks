package main

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

	// Writer captures the latency to execute the INSERT write query
	writerHists []*hdrhistogram.Histogram
	// Read Select captures the latency to execute the SELECT+DELETE+INSERT read combination
	readerReadHists []*hdrhistogram.Histogram
	// E2E captures the end-to-end latency, `read_time - write_time`
	readerE2EHists []*hdrhistogram.Histogram
}

func NewMetrics(writers, readers int) *Metrics {
	m := &Metrics{
		writerHists:     make([]*hdrhistogram.Histogram, writers),
		readerReadHists: make([]*hdrhistogram.Histogram, readers),
		readerE2EHists:  make([]*hdrhistogram.Histogram, readers),
	}
	for i := range m.writerHists {
		m.readerE2EHists[i] = hdrhistogram.New(1, int64(600e9), 3)
	}
	for i := range m.readerReadHists {
		m.readerE2EHists[i] = hdrhistogram.New(1, int64(600e9), 3)
	}
	for i := range m.readerE2EHists {
		m.readerE2EHists[i] = hdrhistogram.New(1, int64(600e9), 3)
	}
	return m
}
