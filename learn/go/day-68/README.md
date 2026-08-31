# Day 68 — Linting with golangci-lint

**Phase:** Advanced Testing & Quality (Days 66–70)

Shared linter config and fixes for errcheck findings in production code.

## What changed from Day 67

| Area | Change |
|------|--------|
| `.golangci.yml` | `govet`, `staticcheck`, `errcheck`, `ineffassign` |
| `internal/db/tx.go` | Explicit rollback in defer |
| `internal/httpapi/response.go` | Ignore encode error after headers sent |
| `scripts/run-lint.ps1` | One command to lint the module |

## Install & run

```powershell
cd learn/go/day-68
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2
.\scripts\run-lint.ps1
```

Or directly:

```powershell
golangci-lint run ./...
```

## Tests

```powershell
go test ./...
go test -tags=integration ./internal/httpapi/...
```

Generated protobuf code under `internal/gen` is excluded from lint rules.

## Local development

Same as Day 67:

```powershell
Copy-Item .env.example .env
go run ./cmd/api
```
