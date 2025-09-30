package metrics

import (
	"fmt"
	"net/http"

	"github.com/aliskhannn/gc-monitor/internal/model"
)

// service defines the interface for retrieving Metrics.
type service interface {
	// GetMetrics retrieves the current Metrics from the Collector.
	GetMetrics() model.Metrics
}

// Handler serves HTTP endpoints for Prometheus metrics.
type Handler struct {
	service service
}

// NewHandler creates a new Metrics HTTP handler with the given service.
func NewHandler(s service) *Handler {
	return &Handler{service: s}
}

// ServeHTTP writes the current metrics in Prometheus format to the response.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	metrics := h.service.GetMetrics()

	fmt.Fprintf(w, "gc_monitor_alloc_bytes %d\n", metrics.Alloc)
	fmt.Fprintf(w, "gc_monitor_total_alloc_bytes %d\n", metrics.TotalAlloc)
	fmt.Fprintf(w, "gc_monitor_sys_bytes %d\n", metrics.Sys)
	fmt.Fprintf(w, "gc_monitor_num_gc %d\n", metrics.NumGC)
	fmt.Fprintf(w, "gc_monitor_last_gc_pause_ns %d\n", metrics.LastGCPause.Nanoseconds())
}
