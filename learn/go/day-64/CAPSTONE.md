# Architecture Capstone — Day 61

**Phase:** Project Layout & Architecture (Days 56–60)

Refactor milestone complete — external behavior unchanged, structure improved.

## Phase recap

| Day | Focus | Key deliverable |
|-----|-------|-----------------|
| 56 | Standard layout | `domain/` → `service/` → `httpapi/` + `repository/` |
| 57 | Clean architecture | Domain rules; service orchestration; repo = persistence only |
| 58 | Dependency injection | `app.Wire`, fakes, inject `clock.Clock` |
| 59 | Rich domain | Value objects, validation in domain, thin handlers |
| 60 | **Capstone** | All of the above integrated, tested, documented |

## Capstone checklist

### Structure

- [x] `cmd/api/main.go` is thin: config → `app.Wire` → server
- [x] `internal/app/wire.go` is the composition root
- [x] No business logic in `httpapi/` handlers
- [x] No ownership/policy checks in `repository/sqlite/`
- [x] Deleted stubs: `bookmark_usecase.go`, `bookmark_validated.go`

### Domain (Day 57 + 59)

- [x] `CanModifyBookmark`, `CanBulkCreate` implemented + tested
- [x] `NewTitle`, `NewBookmarkURL`, `ValidateCreateInput` implemented + tested
- [x] `domain.IsValidation` mapped to HTTP 400 in `httpapi/errors.go`

### Services (Day 57 + 58)

- [x] Update/Delete: Get → domain rule → repo
- [x] `fake.Bookmarks` implemented; service tests active
- [x] `clock.Clock` injected in JWT + memory store

### Transport

- [x] Handlers: parse → service → render only
- [x] `middleware.InjectDeps` removed
- [x] Domain errors mapped in `writeDomainError`

### Tests

- [x] `go test ./internal/domain/...`
- [x] `go test ./internal/service/...`
- [x] `go test ./internal/httpapi/...`
- [x] `go test ./internal/repository/sqlite/...`
- [x] `go test ./...`

### Architecture guards (Day 61)

- [x] `internal/architecture/layers_test.go`
- [x] `scripts/verify-layers.ps1`
- [x] `PACKAGE_DIAGRAM.md`

### Docs

- [x] `README.md` — package map, test pyramid, run instructions
- [x] `ARCHITECTURE.md`, `DI.md`, `DOMAIN.md`
- [ ] `REFLECTION.md` — fill in when ready

## Verify

```powershell
cd learn/go/day-61
go test ./...
go build ./cmd/api
.\scripts\verify-layers.ps1
go run ./cmd/api
```
