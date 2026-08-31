# Day 67 — Test Fixtures and Containers

**Phase:** Advanced Testing & Quality (Days 66–70)

Centralized test environment setup with reliable teardown and optional Docker-based Postgres via testcontainers.

## What changed from Day 66

| Area | Package | Change |
|------|---------|--------|
| Test env | `internal/test/env` | `SetupHTTP` centralizes DB + httptest wiring |
| Parallel safety | `internal/test/env/parallel_test.go` | Isolated temp DB per parallel subtest |
| Containers | `internal/test/env/postgres_container.go` | Optional Postgres via testcontainers (`integration` tag) |

## Run tests

```powershell
cd learn/go/day-67

# Default suite
go test ./...

# Slow HTTP integration tests
go test -tags=integration ./internal/httpapi/...

# Postgres container lifecycle (requires Docker)
go test -tags=integration ./internal/test/env/...
```

## Patterns

- **Setup/teardown:** `env.SetupHTTP` registers `t.Cleanup` for server, repositories, and DB — runs even when tests fail.
- **Parallel tests:** use `t.Parallel()` with `t.TempDir()`-backed SQLite files so workers do not share state.
- **Containers:** `env.PostgresContainer` demonstrates real service startup; the app still uses SQLite — Postgres tests validate container lifecycle only.

## Local development

Same as Day 66:

```powershell
Copy-Item .env.example .env
go run ./cmd/api
```
