# Performance — Days 86–90

## Day 86 — pprof

Debug port `PPROF_PORT` (default `6060`). See `./scripts/capture-profile.ps1`.

## Day 87 — Allocations

- `internal/perf/buffer` — `sync.Pool` for JSON encoding
- Preallocated slices in hot paths
- Track with `go test -benchmem`

## Day 88 — Concurrency

- `internal/concurrency` — bounded `errgroup` for outbox relay
- `OUTBOX_WORKERS` env (default 4)
- Profile goroutines: `/debug/pprof/goroutine`

## Day 89 — DB & HTTP

- `db.ExplainQueryPlan` for slow SQL
- `internal/httpclient` — keep-alive + timeouts

## Day 90 — Evidence

See [PERF_RESULTS.md](./PERF_RESULTS.md) and `./scripts/verify-performance.ps1`.
