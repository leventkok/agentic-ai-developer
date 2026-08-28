# Day 50 — Databases (II) Capstone

**Phase:** Databases (II) & Repositories (Days 46–50)

## Run
```powershell
cd learn/go/day-50
go mod tidy
go run ./cmd/api
go test ./...
```

## Phase summary

| Day | Topic | Key deliverable |
|-----|-------|-----------------|
| 46 | Repository pattern | `repository.Bookmarks` interface + sqlite/memory |
| 47 | Query organization | `internal/db/queries.go` |
| 48 | DB testing | `testutil.OpenTestDB`, integration tests |
| 49 | Connection pooling | Pool config in env + shutdown logging |
| 50 | Review | `DATA_LAYER.md`, full test coverage |

## Structure

```
internal/repository/        ← interface (domain ops, no SQL)
internal/repository/sqlite/ ← production impl
internal/repository/memory/ ← test fake
internal/db/queries.go      ← all SQL in one place
internal/db/testutil/       ← integration test helpers
```

See `DATA_LAYER.md` for full documentation.

## Try it

```powershell
go run ./cmd/api

curl.exe http://localhost:8080/bookmarks
curl.exe -X POST http://localhost:8080/bookmarks/bulk `
  -H "Content-Type: application/json" `
  -d '{"bookmarks":[{"title":"X","url":"https://x.dev"}]}'
```

Data persists in `bookmarks.db`. Pool stats log on Ctrl+C shutdown.
