# Day 46 — Repository Pattern

**Phase:** Databases (II) & Repositories (Days 46–50)

## Focus today
- `repository.Bookmarks` interface — domain ops, no SQL
- `repository/sqlite` — production implementation
- `repository/memory` — fake for unit tests
- Handlers depend on interface, injected via constructor

## Run
```powershell
cd learn/go/day-46
go mod tidy
go test ./...
go run ./cmd/api
```

See `DATA_LAYER.md` for full architecture (days 46–50 build on each other).
