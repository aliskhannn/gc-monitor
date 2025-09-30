package metrics

import "github.com/aliskhannn/gc-monitor/internal/model"

// Collector defines an interface for collecting Metrics.
type Collector interface {
	// Collect reads memory statistics from the runtime and returns
	// them as a Metrics struct.
	Collect() model.Metrics
}

// Service provides access to collected Metrics using a Collector.
type Service struct {
	collector Collector
}

// NewService creates a new Service instance with the given Collector.
func NewService(c Collector) *Service {
	return &Service{collector: c}
}

// GetMetrics retrieves the current Metrics from the Collector.
func (s *Service) GetMetrics() model.Metrics {
	return s.collector.Collect()
}
