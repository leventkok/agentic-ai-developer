# Day 90 — Performance Capstone

## Checklist

- [x] pprof on debug port (`PPROF_PORT`)
- [x] Allocation tuning (`sync.Pool`, preallocated slices)
- [x] Bounded concurrency for outbox relay (`OUTBOX_WORKERS`)
- [x] EXPLAIN query plans + HTTP keep-alive client
- [x] Documented before/after — [PERF_RESULTS.md](./PERF_RESULTS.md)
- [x] Benchmark regression guard

## Verify

```powershell
./scripts/verify-performance.ps1
./scripts/load-test.ps1
./scripts/capture-profile.ps1
```

## Next phase

Days 91–95: Team practices (code review, ADRs, onboarding docs).
