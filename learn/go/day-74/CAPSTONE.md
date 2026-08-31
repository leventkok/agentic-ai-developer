# Quality Capstone — Day 70

**Phase:** Advanced Testing & Quality (Days 66–70)

Automated quality suite: integration tests, golangci-lint, coverage gates, and CI workflow template.

## Phase recap

| Day | Focus | Key deliverable |
|-----|-------|-----------------|
| 66 | Integration testing | `httptest.Server` E2E against real SQLite |
| 67 | Fixtures & containers | `internal/test/env`, optional testcontainers |
| 68 | Linting | `.golangci.yml` with core analyzers |
| 69 | Coverage | Coverage profile + minimum threshold |
| 70 | **Capstone** | Full quality gate script + hardened test suite |

## Capstone checklist

### Integration tests

- [x] Auth register/login E2E
- [x] Bookmark create → update → delete E2E
- [x] Forbidden update when not owner (403)
- [x] Parallel-safe HTTP tests
- [x] Slow tests behind `integration` build tag

### Linting (Day 68)

- [x] `.golangci.yml` — `govet`, `staticcheck`, `errcheck`, `ineffassign`
- [x] Generated code excluded (`internal/gen`)
- [x] Real fixes in `internal/db/tx.go`, `internal/httpapi/response.go`

### Coverage & gates (Day 69)

- [x] `scripts/verify-quality.ps1` — tests + coverage threshold (55%) + lint
- [x] Extra service tests for create/list paths

### CI template (Day 70)

- [x] `.github/workflows/ci.yml` — reference workflow (copy to repo root if needed)

### Docs

- [x] README explains unit vs integration vs quality commands
- [ ] `REFLECTION.md` — fill when ready

## Verify

```powershell
cd learn/go/day-74
go test ./...
go test -tags=integration ./internal/httpapi/...
.\scripts\verify-quality.ps1
```

## Test pyramid (this project)

| Layer | Command | Purpose |
|-------|---------|---------|
| Unit | `go test ./internal/domain/... ./internal/service/...` | Fast rules and orchestration |
| Repository | `go test ./internal/repository/sqlite/...` | SQL + migrations |
| Integration | `go test ./internal/httpapi/...` | HTTP wiring end-to-end |
| Slow / Docker | `go test -tags=integration ./internal/test/env/...` | Containers optional |
| Quality gate | `.\scripts\verify-quality.ps1` | Lint + coverage threshold |

## What you built

A credible regression suite — the same ingredients teams use before merge: meaningful integration tests, shared linter config, and measurable coverage on critical packages.
