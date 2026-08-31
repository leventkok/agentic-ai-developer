# Migrations

## Apply (automatic)

Migrations run on API startup:

```go
db.RunMigrations(database, "migrations")
```

## Files

| Version | Description |
|---------|-------------|
| 001 | Create `bookmarks` table |
| 002 | Index on `url` |
| 003 | Seed dev data |
| 004 | `bookmark_audit` table |

## Rollback (dev only)

```go
db.RollbackMigration(database, "migrations") // reverts latest version
```

Run from a small Go snippet or test — not exposed via HTTP.

## Fresh start

Delete `bookmarks.db` and restart the API to re-apply all migrations from scratch.
