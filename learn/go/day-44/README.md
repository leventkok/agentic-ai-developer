# Day 44 — Transactions

**Phase:** Databases (I) (Days 41–45)

## Run
```powershell
cd learn/go/day-44
go mod tidy
go run ./cmd/api
go test ./...
```

## What's new vs Day 43

| Day 43 | Day 44 |
|--------|--------|
| Single-step writes | Multi-step writes in transactions |
| No audit trail | `bookmark_audit` child table |
| One insert at a time | `BulkCreate` — all or nothing |

## Key files

| File | Purpose |
|------|---------|
| `internal/db/tx.go` | `RunInTx` — begin, defer rollback, commit |
| `internal/store/db.go` | `BulkCreate`, transactional `Update` + audit |
| `migrations/004_*` | `bookmark_audit` table |

## Transaction pattern

```go
err := db.RunInTx(ctx, database, func(tx *sql.Tx) error {
    // step 1
    // step 2 — if this fails, step 1 is rolled back
    return nil
})
```

`defer tx.Rollback()` runs on every path; after successful `Commit`, rollback is a no-op.

## Try it

```powershell
go run ./cmd/api

curl.exe -X POST http://localhost:8080/bookmarks/bulk `
  -H "Content-Type: application/json" `
  -d '{"bookmarks":[{"title":"A","url":"https://a.dev"},{"title":"B","url":"https://b.dev"}]}'

curl.exe -X PATCH http://localhost:8080/bookmarks/1 `
  -H "Content-Type: application/json" `
  -d '{"title":"Updated title"}'
```

Update writes both the bookmark row and an audit row atomically.
