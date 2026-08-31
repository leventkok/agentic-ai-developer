# Observability Capstone — Day 75

**Phase:** Observability & Resilience (Days 71–75)

Operable MVP with structured logs, Prometheus metrics, OpenTelemetry traces, and resilient reads.

## Phase recap

| Day | Focus | Key deliverable |
|-----|-------|-----------------|
| 71 | Structured logging | JSON `slog` with request fields |
| 72 | Metrics | `/metrics` counters + histograms |
| 73 | Tracing | OTel HTTP spans (stdout) |
| 74 | Resilience | Retry + circuit breaker on reads |
| 75 | **Capstone** | Full stack + runbook + verification |

## Capstone checklist

- [x] Structured logs: request_id, method, path, status, duration
- [x] Prometheus `/metrics` scrape endpoint
- [x] OpenTelemetry tracing initialized in `cmd/api`
- [x] Retry + breaker on idempotent repository reads
- [x] `RUNBOOK.md` for on-call steps
- [x] `scripts/verify-observability.ps1`
- [x] `go test ./...` green

## Verify

```powershell
cd learn/go/day-76
go test ./...
.\scripts\verify-observability.ps1
go run ./cmd/api
```

## Dashboard sketch (what to monitor)

| Panel | Signal |
|-------|--------|
| Traffic | `http_requests_total` rate |
| Errors | 5xx ratio from logs/metrics |
| Latency | `http_request_duration_seconds` p95 |
| Saturation | DB pool / goroutine count (future) |
| Breaker | gobreaker state transitions (logs) |

## What you built

A service you can **operate** — not only deploy. Logs tell you what happened, metrics show trends, traces show where time went, and breakers stop hammering sick dependencies.
