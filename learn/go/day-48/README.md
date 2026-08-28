# Day 48 — Testing with Databases

**Phase:** Databases (II) & Repositories (Days 46–50)

## Focus today
- `internal/db/testutil` — `OpenTestDB`, `ResetTables`, `MigrationsDir`
- Integration tests in `repository/sqlite/store_integration_test.go`
- Each test gets isolated temp DB; `ResetTables` clears data between cases

## Run
```powershell
cd learn/go/day-48
go mod tidy
go test ./internal/repository/sqlite/... -v
go test ./...
```

See `DATA_LAYER.md` for test instructions.
