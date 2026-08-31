# Clean Architecture — Day 57

## Dependency rule

Dependencies point **inward** only: `httpapi` → `service` → `domain` ← `repository`.

## Checklist

- [x] `internal/domain/bookmark_rules.go` — ownership and bulk-create rules
- [x] `internal/service/bookmark.go` — Update/Delete orchestration
- [x] `internal/repository/sqlite/store.go` — persistence only, no policy
- [x] `internal/repository/memory/store.go` — persistence only
- [x] Domain and service tests green

## Verify

```powershell
rg "net/http|database/sql" internal/domain/
go test ./internal/domain/... ./internal/service/...
```
