# Domain Models & Services — Day 59

## Checklist

- [x] `domain/title.go`, `domain/url.go` — value objects with constructors
- [x] `domain/bookmark_validate.go` — create/update/bulk validation
- [x] `domain/value_objects_test.go` — table-driven tests
- [x] `service/bookmark.go` — validates via domain before repo
- [x] `httpapi/errors.go` — `IsValidation` → 400
- [x] `httpapi/bookmarks.go` — no business validation calls
- [x] `validation/bookmark.go` — `ParseID` only

## Verify

```powershell
rg "validation.Validate" internal/httpapi/
# expect: no matches
go test ./internal/domain/...
```
