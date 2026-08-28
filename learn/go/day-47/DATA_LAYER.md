# Data Layer — Day 50

## Architecture

```
HTTP Handler
    ↓ repository.Bookmarks (interface)
SQLite Store / Memory Fake
    ↓ db/queries.go (SQL constants)
database/sql + SQLite
```

Handlers never import SQL. All persistence goes through `repository.Bookmarks`.

## Packages

| Package | Role |
|---------|------|
| `internal/repository` | Domain interface + `ErrNotFound` |
| `internal/repository/sqlite` | Production SQLite implementation |
| `internal/repository/memory` | In-memory fake for unit tests |
| `internal/db/queries.go` | Named SQL constants |
| `internal/db/pool.go` | Connection pool config + stats |
| `internal/db/testutil` | Test DB setup + table reset |

## Migrations

See `MIGRATIONS.md`. Applied on startup via `db.RunMigrations`.

## Pool settings

| Env | Default | Notes |
|-----|---------|-------|
| `DB_PATH` | `bookmarks.db` | SQLite file path |
| `DB_MAX_OPEN_CONNS` | `1` | SQLite single-writer |
| `DB_MAX_IDLE_CONNS` | `1` | Idle connections kept |
| `DB_CONN_MAX_LIFETIME_MIN` | `0` | 0 = no limit |

Pool stats logged on graceful shutdown.

## Tests

```powershell
# Unit tests (memory fake)
go test ./internal/handler/... -v

# Integration tests (real SQLite, isolated temp DB per test)
go test ./internal/repository/sqlite/... -v

# Full suite
go test ./...
```

Integration tests use `testutil.OpenTestDB` + `testutil.ResetTables` for isolation.

## Transactions

- `Update` — bookmark row + audit row atomically
- `BulkCreate` — all inserts or none
- `db.RunInTx` — generic helper with `defer Rollback()`
