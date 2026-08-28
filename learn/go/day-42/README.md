# Day 42 — Migrations and Schema Design

**Phase:** Databases (I) (Days 41–45)

## Run
```powershell
cd learn/go/day-42
go mod tidy
go run ./cmd/api
go test ./internal/db/... -v
```

## What's new vs Day 41

| Day 41 | Day 42 |
|--------|--------|
| Single SQL file | Versioned up/down migrations |
| No tracking | `schema_migrations` table |
| Basic columns | `created_at`, `updated_at`, index on `url` |

## Migration files

```
migrations/
  001_create_bookmarks.up.sql
  001_create_bookmarks.down.sql
  002_add_url_index.up.sql
  002_add_url_index.down.sql
  003_seed_bookmarks.up.sql
  003_seed_bookmarks.down.sql
```

## API

```go
db.RunMigrations(database, "migrations")  // apply pending
db.RollbackMigration(database, "migrations") // revert latest (dev)
```

See `SCHEMA.md` for table design.
