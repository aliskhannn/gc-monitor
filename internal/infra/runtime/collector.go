package runtime

import (
	"runtime"
	"time"

	"github.com/aliskhannn/gc-monitor/internal/model"
)

// Collector is responsible for collecting memory and GC metrics
// from the Go runtime.
type Collector struct {
}

// NewCollector creates a new instance of Collector.
func NewCollector() *Collector {
	return &Collector{}
}

// Collect reads memory statistics from the runtime and returns
// them as a Metrics struct.
func (c *Collector) Collect() model.Metrics {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	lastPause := time.Duration(0)
	if memStats.NumGC > 0 {
		// The PauseNs array is circular, use modulo 256 to get the last pause.
		lastPause = time.Duration(memStats.PauseNs[(memStats.NumGC+255)%256])
	}

	return model.Metrics{
		Alloc:       memStats.Alloc,
		TotalAlloc:  memStats.TotalAlloc,
		Sys:         memStats.Sys,
		NumGC:       memStats.NumGC,
		LastGCPause: lastPause,
	}
}
