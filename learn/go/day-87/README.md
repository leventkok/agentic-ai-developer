# Day 87 — Memory and Allocation Tuning

**Phase:** Performance (Days 86–90)

Reduce allocations in hot paths using `sync.Pool` and preallocated slices.

## Changes

| Area | Optimization |
|------|--------------|
| `internal/perf/buffer` | Pooled `bytes.Buffer` for JSON responses |
| `sqlite.Store.List` | Preallocated slice capacity |
| Benchmarks | `-benchmem` via `go test -bench=. -benchmem ./internal/repository/sqlite/...` |

Also includes Day 86 pprof (`PPROF_PORT=6060`).

```powershell
cd learn/go/day-87
go test ./...
go test -bench=BenchmarkStoreList -benchmem ./internal/repository/sqlite/...
```
