# Domain Models & Services — Day 59

## Anemic vs rich models

**Before (anemic):** plain structs + validation in `internal/validation/`

```go
type CreateBookmarkInput struct {
    Title string  // anything goes until validation package runs
    URL   string
}
```

**After (rich):** value objects enforce invariants at construction

```go
title, err := domain.NewTitle(raw)       // rejects empty / too long
url, err := domain.NewBookmarkURL(raw)   // rejects invalid schemes
bookmark, err := domain.NewBookmark(in)  // never invalid if err == nil
```

Invalid state cannot exist inside `Title` or `BookmarkURL` — the type guarantees it.

## Error flow

```
Handler          Service              Domain
  │                 │                    │
  │ parse JSON      │                    │
  │────────────────>│ NewBookmark(in)    │
  │                 │───────────────────>│ validate invariants
  │                 │<── ValidationError ─│
  │<── 400 mapped ──│                    │
```

Domain returns **typed errors** (`ValidationError`, `ErrForbidden`, …).  
Transport maps them — handlers never branch on string messages.

## Today's refactor targets

- [ ] **`domain/title.go`, `domain/url.go`** — value objects with constructors
- [ ] **`domain/bookmark_factory.go`** — `NewBookmark` validates before persistence
- [ ] **`domain/validation_errors.go`** — typed validation errors + `IsValidation`
- [ ] **`service/bookmark.go`** — Create/Update call domain constructors first
- [ ] **`httpapi/errors.go`** — map `domain.IsValidation(err)` → 400 Bad Request
- [ ] **`httpapi/bookmarks.go`** — thin: parse JSON → service → write response (no validation pkg)
- [ ] **`internal/validation/bookmark.go`** — delete or reduce to HTTP-only concerns (ParseID)
- [ ] **`domain/value_objects_test.go`** — un-skip table tests

## Handler thinness checklist

A handler should only:

1. Extract auth user from context
2. Decode request body / path params
3. Call one service method
4. Map error → status via `writeDomainError`
5. Write JSON response

No business rules. No URL parsing logic. No ownership checks.

## Verify

```powershell
cd learn/go/day-59
go test ./internal/domain/...
go test ./...
```

```powershell
# Validation logic should live in domain, not validation package
rg "requireURL|requireNonEmpty" internal/validation/
# after refactor: only ParseID / HTTP-specific helpers remain
```
