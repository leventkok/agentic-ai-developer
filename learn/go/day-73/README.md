# Day 73 — OpenTelemetry Tracing

**Phase:** Observability & Resilience (Days 71–75)

HTTP spans via OpenTelemetry with stdout exporter for local learning.

## What changed from Day 72

| Area | Package | Change |
|------|---------|--------|
| Tracing | `internal/observability/tracing` | OTel tracer provider + stdout export |
| Middleware | `internal/middleware/tracing.go` | `otelhttp` wraps each request |
| Startup | `cmd/api/main.go` | Initializes tracing on boot |

## Run

```powershell
cd learn/go/day-73
go test ./...
go run ./cmd/api
# Trace output prints to stdout when requests hit the server
```

Metrics remain on `/metrics` from Day 72.
