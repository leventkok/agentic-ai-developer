# Day 72 — Prometheus Metrics

**Phase:** Observability & Resilience (Days 71–75)

`/metrics` endpoint with request counters and latency histograms.

## What changed from Day 71

| Area | Package | Change |
|------|---------|--------|
| Metrics | `internal/observability/metrics` | Prometheus counters + histograms |
| Middleware | `internal/middleware/metrics.go` | Records each HTTP request |
| Router | `GET /metrics` | Scrape endpoint |

## Verify

```powershell
cd learn/go/day-72
go test ./...
go run ./cmd/api
curl http://localhost:8080/metrics
```

Routes are normalized to `/bookmarks/{id}` labels to limit cardinality.
