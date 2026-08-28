# Day 47 — Query Organization

**Phase:** Databases (II) & Repositories (Days 46–50)

## Focus today
- All SQL in `internal/db/queries.go`
- Named constants: `SQLListBookmarks`, `SQLGetBookmarkByID`, etc.
- SQLite store references constants — no inline SQL strings

## Run
```powershell
cd learn/go/day-47
go mod tidy
go test ./...
```

See `internal/db/queries.go` and `DATA_LAYER.md`.
