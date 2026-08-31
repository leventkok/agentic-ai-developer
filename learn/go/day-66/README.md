# Day 66 — Integration Testing

**Phase:** Advanced Testing & Quality (Days 66–70)

HTTP integration tests against a real SQLite test database using `httptest.Server` and shared fixtures.

## What changed from Day 65

| Area | Package | Change |
|------|---------|--------|
| Fixtures | `internal/test/fixtures` | `RegisterUser` helper for auth setup |
| Integration | `internal/httpapi/integration_test.go` | E2E auth + bookmark CRUD via HTTP |
| Slow tests | `internal/httpapi/integration_slow_test.go` | Build tag `integration` for heavier paths |

## Run tests

```powershell
cd learn/go/day-66

# Default suite (includes integration tests in httpapi)
go test ./...

# Optional slow integration tests
go test -tags=integration ./internal/httpapi/...
```

## Integration test pattern

1. Open ephemeral SQLite DB via `testutil.OpenTestDB`
2. Wire real repositories + `httpapi.NewRouter`
3. Start `httptest.NewServer`
4. Call real HTTP endpoints with `http.DefaultClient`
5. Seed users with `fixtures.RegisterUser`

Slow or long-running tests live behind `//go:build integration` so normal `go test ./...` stays fast.

## Local development

Same as Day 65 — combined HTTP + gRPC server:

```powershell
Copy-Item .env.example .env
go run ./cmd/api
```

See `README` sections in prior days for protobuf generation and gRPC client usage.
