# Day 58 — Dependency Injection (Complete)

**Phase:** Project Layout & Architecture (Days 56–60)

Builds on Day 57. Explicit wiring, interface boundaries, fakes for service tests.

## What changed

- `internal/app/wire.go` — composition root
- `internal/clock` — injectable time in JWT and memory store
- `internal/service/testing/fake` — in-memory repos for unit tests
- `middleware.InjectDeps` removed; thin `cmd/api/main.go`

## Run

```powershell
cd learn/go/day-58
go test ./internal/service/...
go test ./...
go run ./cmd/api
```

## Next

Day 59 — rich domain models (`learn/go/day-59/`).
