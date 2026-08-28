# Retrospective — Day 45 (Databases I capstone)

## What went well
- Full CRUD persists in SQLite — API survives restarts
- Migrations are versioned, tracked, and reversible
- Prepared statements prevent SQL injection
- Transactions keep bookmark updates and audit rows consistent
- Integration smoke test proves handler → store → DB path works

## What was hard
- SQLite datetime format vs Go `time.Time` parsing
- Knowing when to use transactions vs single statements
- Migration rollback order (latest first)
- Keeping memory store in sync for unit tests while DB is production path

## Production patterns I'll reuse
- **`defer tx.Rollback()`** after every `BeginTx`
- **Repository interface** so handlers never import SQL
- **Migrations in repo** — schema is code
- **Integration tests** against real DB, not just mocks
- **Parameterized queries** — always `?`, never string concat

## Next phase (Days 46–50)
- Formal repository package separation
- Centralized SQL query constants
- DB test helpers with table reset
- Connection pool tuning for production
