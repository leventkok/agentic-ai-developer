# Day 45 — Databases (I) Capstone

**Phase:** Databases (I) (Days 41–45)

## Run
```powershell
cd learn/go/day-45
go mod tidy
go run ./cmd/api
go test ./internal/handler/... -run Smoke -v
```

## Milestone

Your Bookmarks API is now **fully persistent**:
- SQLite via `database/sql`
- Versioned migrations (`migrations/*.up.sql`)
- Prepared statements + parameterized queries
- Transactions for multi-step writes (update + audit, bulk create)

Data survives process restarts in `bookmarks.db`.

## Migrations

```powershell
# Applied automatically on startup via db.RunMigrations(database, "migrations")
go run ./cmd/api
```

See `MIGRATIONS.md` for manual rollback during development.

## Smoke test

Integration test exercises full CRUD against a real SQLite file:

```powershell
go test ./internal/handler/... -run TestCRUDSmoke -v
```

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /bookmarks | List all |
| POST | /bookmarks | Create one |
| POST | /bookmarks/bulk | Create many (transaction) |
| GET | /bookmarks/{id} | Get by ID |
| PATCH | /bookmarks/{id} | Update (+ audit row) |
| DELETE | /bookmarks/{id} | Delete |
