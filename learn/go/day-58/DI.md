# Dependency Injection — Day 58

## Checklist

- [x] `internal/app/wire.go` — single composition root
- [x] `internal/clock/clock.go` — injected into JWT + memory store
- [x] `internal/service/testing/fake/` — fake bookmarks and auth
- [x] Service tests with fakes (no DB)
- [x] `middleware.InjectDeps` removed from `DefaultStack`
- [x] `cmd/api/main.go` calls `app.Wire`

## Verify

```powershell
go test ./internal/service/...
rg "InjectDeps" internal/
# expect: no matches
```
