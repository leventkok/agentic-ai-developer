# Performance Results — Day 90

Measured on empty SQLite DB (local dev). Use as baseline for future regressions.

| Metric | Before (Day 85) | After (Day 90) | Change |
|--------|-----------------|----------------|--------|
| `BenchmarkStoreList` ns/op | ~3500 | ~2800 | preallocated slice + pooled JSON |
| `BenchmarkStoreList` allocs/op | ~15 | ~12 | buffer pool on HTTP encode |
| List handler p99 (200 req) | ~8ms | ~6ms | cache-aside hit after warm |

## Optimizations applied

1. **Preallocated slice** in `sqlite.Store.List` (capacity 32)
2. **`sync.Pool` JSON buffers** in `httpapi.writeJSON`
3. **Parallel outbox relay** with bounded `errgroup` workers
4. **HTTP keep-alive client** (`internal/httpclient`) for outbound calls

## How to reproduce

```powershell
go test -bench=BenchmarkStoreList -benchmem ./internal/repository/sqlite/...
./scripts/load-test.ps1 -Requests 200
./scripts/capture-profile.ps1
```

## Regression guard

`TestListBenchmarkRegressionGuard` fails CI if list benchmark exceeds 1ms/op.
