# Retrospective — Day 50 (Databases II capstone)

## What went well
- Clean separation: handlers → repository interface → SQLite
- SQL centralized in `queries.go` — easy to audit
- Integration tests with real DB catch SQL typos mocks miss
- Connection pool tuned via config, closed on shutdown
- Memory fake keeps handler unit tests fast

## What was hard
- Refactoring from `internal/store` to `internal/repository` without breaking tests
- SQLite pool defaults (MaxOpenConns=1) vs Postgres mental model
- Test isolation — each test needs its own DB or table reset
- Mapping DB rows to domain types at the right boundary

## Production patterns I'll reuse
- **Repository interface** — swap SQLite for Postgres without touching handlers
- **Named SQL constants** — no scattered query strings
- **testutil.ResetTables** — fast integration test isolation
- **Pool stats on shutdown** — early signal of connection starvation
- **Constructor injection** — `handler.New(repo)` not global DB

## Next phase (Days 51–55)
- JWT authentication
- Authorization middleware
- RBAC for bookmark ownership
