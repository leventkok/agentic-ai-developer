# Day 86 — Profiling with pprof

**Phase:** Performance (Days 86–90)

CPU and heap profiling via a dedicated debug port. See [PERFORMANCE.md](./PERFORMANCE.md).

```powershell
cd learn/go/day-86
go test ./...
go run ./cmd/api
./scripts/capture-profile.ps1
```
