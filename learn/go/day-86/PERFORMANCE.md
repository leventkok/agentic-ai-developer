# Performance — Day 86

## pprof

The API exposes Go's built-in profiler on a **separate debug port** (default `6060`):

| Endpoint | Purpose |
|----------|---------|
| `/debug/pprof/profile` | CPU profile |
| `/debug/pprof/heap` | Heap allocations |
| `/debug/pprof/goroutine` | Goroutine stacks |

Set `PPROF_PORT=` empty to disable in production.

## Workflow

1. Start API: `go run ./cmd/api`
2. Generate load against `GET /bookmarks`
3. Capture: `./scripts/capture-profile.ps1`
4. Inspect: `go tool pprof -http=:8081 profiles/cpu.prof`

## Benchmarks

```powershell
go test -bench=. -benchmem ./internal/repository/sqlite/...
```

Use benchmarks as baselines before optimizing (Days 87–90).
