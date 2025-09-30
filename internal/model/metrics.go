package model

import "time"

// Metrics holds memory and garbage collection statistics.
type Metrics struct {
	Alloc       uint64        // currently allocated bytes
	TotalAlloc  uint64        // total bytes allocated (even if freed)
	Sys         uint64        // total bytes obtained from the OS
	NumGC       uint32        // number of completed GC cycles
	LastGCPause time.Duration // duration of the last GC pause
}
