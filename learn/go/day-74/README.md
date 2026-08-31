# Day 74 — Retries, Circuit Breakers, and Timeouts

**Phase:** Observability & Resilience (Days 71–75)

Resilience wrapper on bookmark reads with retry + circuit breaker; writes stay direct (non-idempotent).

## What changed from Day 73

| Area | Package | Change |
|------|---------|--------|
| Retry | `internal/resilience/retry.go` | Exponential backoff for idempotent reads |
| Breaker | `internal/repository/resilient` | `gobreaker` around List/Get |
| Wiring | `internal/app/wire.go` | SQLite repo wrapped before services |

## Policy

- **Retry:** List/Get only (idempotent reads)
- **No retry:** Create/Update/Delete (side effects)
- **Breaker:** Opens after repeated failures, half-open after timeout

## Tests

```powershell
cd learn/go/day-74
go test ./internal/repository/resilient/...
go test ./...
```
