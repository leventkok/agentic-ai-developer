# Day 70 — Advanced Testing & Quality Capstone

**Phase:** Advanced Testing & Quality (Days 66–70)

Full quality suite: integration tests, golangci-lint, coverage gates, and CI template.

## What changed from Day 69

| Day | Focus | Deliverable |
|-----|-------|-------------|
| 70 | **Capstone** | Full quality gate + hardened E2E (see `CAPSTONE.md`) |

## Run tests

```powershell
cd learn/go/day-70

# Fast unit + integration (default)
go test ./...

# Slow / Docker optional
go test -tags=integration ./internal/httpapi/... ./internal/test/env/...

# Full quality gate (tests + lint + coverage >= 55%)
.\scripts\verify-quality.ps1
```

## Lint locally

```powershell
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2
golangci-lint run ./...
```

## Coverage

```powershell
go test ./internal/service/... ./internal/domain/... ./internal/repository/sqlite/... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

## CI

`.github/workflows/ci.yml` is a **reference workflow** inside this day folder. GitHub Actions reads workflows from the **repository root** — copy or merge into `.github/workflows/` at repo root when you want automated runs on push.

## Local development

Same as Day 65 — combined HTTP + gRPC:

```powershell
Copy-Item .env.example .env
go run ./cmd/api
```

See `CAPSTONE.md` for the full phase checklist.
