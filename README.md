# GC Monitor

GC Monitor is a simple Go service that exposes runtime memory and garbage collection (GC) metrics through an HTTP endpoint in Prometheus format.

---

## Project Structure

```
.
├── cmd/                     # Application entry points
│   └── server/              # Main server
├── internal/                # Internal application packages
│   ├── api/                 
│   │   └── handlers/        # HTTP handlers for endpoints
│   ├── infra/runtime/       # Collector implementation using Go runtime
│   ├── model/               # Data models (Metrics struct, Collector interface)
│   └── service/             # Business logic / Service layer
├── go.mod
├── go.sum
└── README.md
```
---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/aliskhannn/gc-monitor.git
cd gc-monitor
```

2. Download dependencies:

```bash
go mod tidy
```

3. Run the server:

```bash
go run ./cmd/server
```

The server will start on **`http://localhost:8080`**.

---

## Endpoint

**Metrics endpoint:**

```
GET /metrics
```

**Example Response (Prometheus format):**

```
gc_monitor_alloc_bytes 235192
gc_monitor_total_alloc_bytes 235192
gc_monitor_sys_bytes 8083720
gc_monitor_num_gc 0
gc_monitor_last_gc_pause_ns 0
```

**Metrics Description:**

| Metric                         | Description                                   |
| ------------------------------ | --------------------------------------------- |
| `gc_monitor_alloc_bytes`       | Current allocated memory in bytes             |
| `gc_monitor_total_alloc_bytes` | Total memory allocated over time              |
| `gc_monitor_sys_bytes`         | Total memory obtained from OS                 |
| `gc_monitor_num_gc`            | Number of completed garbage collection cycles |
| `gc_monitor_last_gc_pause_ns`  | Duration of the last GC pause in nanoseconds  |

---

## Profiling

The server includes Go’s `pprof` for profiling:

```
http://localhost:8080/debug/pprof/
```

You can inspect CPU, memory, goroutines, and other runtime statistics.

---

## How it Works

1. The **Collector** in `internal/infra/runtime` reads memory and GC stats using `runtime.ReadMemStats`.
2. The **Service layer** in `internal/service` wraps the collector and provides a unified interface to get metrics.
3. The **HTTP Handler** in `internal/api/handlers` exposes the metrics in Prometheus format at `/metrics`.
4. Prometheus or any other monitoring system can scrape the `/metrics` endpoint periodically.