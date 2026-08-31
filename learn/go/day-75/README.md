# Day 75 — Observability & Resilience Capstone

**Phase:** Observability & Resilience (Days 71–75)

Full operability stack: logs, metrics, traces, resilience, and runbook.

## What changed from Day 74

| Day | Focus | Deliverable |
|-----|-------|-------------|
| 75 | **Capstone** | `CAPSTONE.md`, `RUNBOOK.md`, `verify-observability.ps1` |

## Verify everything

```powershell
cd learn/go/day-75
.\scripts\verify-observability.ps1
go run ./cmd/api
curl http://localhost:8080/metrics
```

## Stack summary

- **Logs:** JSON slog (Day 71)
- **Metrics:** Prometheus `/metrics` (Day 72)
- **Traces:** OTel stdout (Day 73)
- **Resilience:** retry + breaker on reads (Day 74)
- **Ops:** `RUNBOOK.md` (Day 75)

See `CAPSTONE.md` for the full checklist.
