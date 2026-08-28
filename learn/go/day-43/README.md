# Day 43 — Queries and Prepared Statements

**Phase:** Databases (I) (Days 41–45)

## Run
```powershell
cd learn/go/day-43
go mod tidy
go run ./cmd/api
go test ./...
```

## What's new vs Day 42

| Day 42 | Day 43 |
|--------|--------|
| DB read-only demo | Full CRUD via `DBStore` |
| Handlers use memory store | Handlers use SQLite store |
| Ad-hoc queries | Prepared statements (`db.Prepare`) |

## Key file

`internal/store/db.go` — implements `BookmarkRepository` with:

- `listStmt`, `getStmt`, `insertStmt`, `updateStmt`, `deleteStmt`
- Parameter binding (`?`) — never string-concat SQL
- `sql.ErrNoRows` → `store.ErrNotFound`
- Wrapped errors: `fmt.Errorf("get bookmark id=%d: %w", id, err)`

## Try it

```powershell
go run ./cmd/api
# DB seed count: 2

curl.exe -X POST http://localhost:8080/bookmarks `
  -H "Content-Type: application/json" `
  -d '{"title":"Test","url":"https://example.com","tags":["demo"]}'

curl.exe http://localhost:8080/bookmarks
```

Data persists in `bookmarks.db` across restarts.
