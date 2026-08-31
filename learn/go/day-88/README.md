# Day 88 — Concurrency Tuning

**Phase:** Performance (Days 86–90)

Bounded parallel outbox relay with `errgroup` and configurable workers.

## Changes

| Area | Optimization |
|------|--------------|
| `internal/concurrency` | `RunLimited` with `errgroup.SetLimit` |
| `internal/messaging/relay` | Parallel publish with `OUTBOX_WORKERS` (default 4) |

Profile goroutines under load: `/debug/pprof/goroutine`

```powershell
cd learn/go/day-88
go test ./...
```
