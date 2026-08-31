# Day 59 — Domain Models & Services (Complete)

**Phase:** Project Layout & Architecture (Days 56–60)

Builds on Day 58. Value objects enforce invariants; validation lives in domain; handlers stay thin.

## What changed

- `NewTitle`, `NewBookmarkURL`, `ValidateCreateInput` in `internal/domain/`
- Service validates before persistence
- Handlers no longer call `validation.ValidateCreateInput`
- `domain.IsValidation` → HTTP 400 in `httpapi/errors.go`
- `validation/bookmark.go` reduced to `ParseID` only

## Run

```powershell
cd learn/go/day-59
go test ./internal/domain/...
go test ./...
go run ./cmd/api
```

## Next

Day 60 — architecture capstone (`learn/go/day-60/`).
