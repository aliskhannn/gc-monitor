package model

import "time"

type Metrics struct {
	Alloc       uint64
	TotalAlloc  uint64
	Sys         uint64
	NumGC       uint32
	LastGCPause time.Duration
}
