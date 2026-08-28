# Day 41 — database/sql Introduction

**Phase:** Databases (I) (Days 41–45)

## Setup
```powershell
cd learn/go/day-41
go mod tidy
```

## Tasks — which file to edit

| Task | What | File |
|------|------|------|
| **1** | Open DB connection (SQLite) | `internal/db/db.go` |
| **2** | Ping on startup (fail fast) | `internal/db/db.go` |
| **3** | SELECT with QueryContext + Scan | `internal/db/bookmarks.go` |
| **4** | defer rows.Close / db.Close | `internal/db/bookmarks.go`, `cmd/api/main.go` |
| **5** | Wire DB in main | `cmd/api/main.go` |

**Schema (already provided):** `migrations/001_init.sql`

See chat lesson for full code + logic.
