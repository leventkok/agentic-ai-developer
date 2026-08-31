# Day 56 — Standard Go Project Layout

**Phase:** Project Layout & Architecture (Days 56–60)

Refactored from Day 55 into layered packages.

## Layout

```
cmd/api/main.go           ← wire dependencies (~75 lines)

internal/
  domain/                 ← Bookmark, User, errors, inputs
  service/                ← AuthService, BookmarkService
  httpapi/                ← handlers, DTOs, router
  ctxkey/                 ← user in context (breaks import cycles)
  repository/             ← interfaces + sqlite/memory
  middleware/             ← auth, rate limit, logging
  config/, db/, auth/, validation/
```

## Dependency flow

```
httpapi → service → repository → db
            ↓
          domain
```

## Run

```powershell
cd learn/go/day-56
go run ./cmd/api
go test ./...
```

See `LAYOUT.md` for package rules.
