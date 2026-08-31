# Clean Architecture — Days 56–60

> Day 58: DI (`DI.md`). Day 59: rich domain (`DOMAIN.md`). Day 60: capstone integration (`CAPSTONE.md`).

## Dependency rule (non-negotiable)

Dependencies point **inward** only:

```
┌─────────────────────────────────────┐
│  Transport (httpapi, middleware)    │  ← HTTP, JSON, status codes
├─────────────────────────────────────┤
│  Service (use cases)                │  ← orchestration, calls domain + repo
├─────────────────────────────────────┤
│  Domain (rules, entities, errors)    │  ← pure Go, no frameworks
├─────────────────────────────────────┤
│  Repository (interfaces + sqlite)   │  ← persistence only
└─────────────────────────────────────┘
```

**Domain must never import:** `net/http`, `database/sql`, `repository`, `httpapi`, `service`.

## Layer responsibilities

| Layer | Owns | Example |
|-------|------|---------|
| **Domain** | Business rules | `CanModifyBookmark(actor, bookmark)` |
| **Service** | Use case flow | Get bookmark → check rule → update |
| **Transport** | Protocol mapping | Decode JSON → call service → write status |
| **Repository** | Storage | SQL, row mapping — no HTTP status codes |

## Today's refactor targets

1. **`internal/domain/bookmark_rules.go`** — move `canModifyBookmark` out of sqlite
2. **`internal/service/bookmark.go`** — Update/Delete orchestrate before calling repo
3. **`internal/repository/sqlite/store.go`** — remove domain policy; persist only
4. **`internal/httpapi/`** — stays thin; no business rules

## Verify dependency direction

```powershell
# Domain must not import outer layers (manual grep)
rg "net/http|database/sql" internal/domain/

# Should return nothing
```

## Test strategy

- **Domain tests** — pure rules, no DB (`domain/bookmark_rules_test.go`)
- **Service tests** — fake repository, no HTTP
- **httpapi tests** — HTTP mapping only (status codes, JSON)
