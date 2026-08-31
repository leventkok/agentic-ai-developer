# Day 57 — Clean Architecture Layers (Complete)

**Phase:** Project Layout & Architecture (Days 56–60)

Domain rules live in `internal/domain/`. Services orchestrate use cases. Repositories persist only.

## What changed

- `CanModifyBookmark`, `CanBulkCreate` in domain with tests
- `BookmarkService` Update/Delete: Get → domain rule → repo
- Ownership checks removed from `repository/sqlite` and `repository/memory`

## Run

```powershell
cd learn/go/day-57
go test ./...
go run ./cmd/api
```

## Next

Day 58 — dependency injection (`learn/go/day-58/`).
