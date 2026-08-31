# Day 60 — Architecture Capstone (Complete)

**Phase:** Project Layout & Architecture (Days 56–60)

Layered bookmarks API with auth, RBAC, rate limiting, domain rules, DI, and rich domain validation.

## Package map

| Package | Purpose |
|---------|---------|
| `cmd/api` | Thin entrypoint — config, server lifecycle |
| `internal/app` | Composition root (`Wire`) — only place that imports sqlite |
| `internal/domain` | Entities, value objects, business rules, typed errors |
| `internal/service` | Use-case orchestration (validate → rule → repo) |
| `internal/httpapi` | HTTP handlers, DTOs, error mapping |
| `internal/repository` | Persistence interfaces |
| `internal/repository/sqlite` | SQLite implementation |
| `internal/repository/memory` | In-memory store (tests) |
| `internal/service/testing/fake` | Fake repos for service unit tests |
| `internal/middleware` | Auth, logging, rate limit, recovery |
| `internal/clock` | Injectable time source |
| `internal/architecture` | Import boundary tests |

## Test pyramid

```
domain/          pure rules, value objects (fastest)
service/         fakes, no HTTP or DB
httpapi/         handler mapping
repository/sqlite/  integration with real DB
architecture/    layer dependency guards
```

## Run

```powershell
cd learn/go/day-60
go test ./...
go build ./cmd/api
.\scripts\verify-layers.ps1
go run ./cmd/api
```

Copy `.env.example` to `.env` and set `JWT_SECRET`.

## Add a new endpoint (pattern)

1. **Domain** — add rule or validation in `internal/domain/`
2. **Service** — orchestrate in `internal/service/`
3. **Handler** — parse JSON → call service → `writeDomainError` in `internal/httpapi/`
4. **Router** — register route in `internal/httpapi/router.go`
5. **Tests** — domain table test → service fake test → httpapi handler test

## Phase folders

| Day | Focus | Status |
|-----|-------|--------|
| 56 | Standard layout | `learn/go/day-56` |
| 57 | Clean architecture | integrated in day-60 |
| 58 | Dependency injection | integrated in day-60 |
| 59 | Rich domain models | integrated in day-60 |
| 60 | Capstone | **this folder** |

See also: `ARCHITECTURE.md`, `DI.md`, `DOMAIN.md`, `CAPSTONE.md`, `PACKAGE_DIAGRAM.md`.
