package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/aliskhannn/gc-monitor/internal/api/handlers/metrics"
	"github.com/aliskhannn/gc-monitor/internal/infra/runtime"
	servicemetrics "github.com/aliskhannn/gc-monitor/internal/service/metrics"
)

func main() {
	// Initialize runtime collector for memory and GC metrics.
	collector := runtime.NewCollector()

	// Initialize service and handler for metrics endpoint.
	service := servicemetrics.NewService(collector)
	handler := metrics.NewHandler(service)

	// Register the metrics endpoint.
	http.Handle("/metrics", handler)

	// Start HTTP server with pprof enabled.
	log.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
