# Day 69 — Code Coverage and Quality Gates

**Phase:** Advanced Testing & Quality (Days 66–70)

Coverage measurement and minimum thresholds on critical packages (service + sqlite).

## What changed from Day 68

| Area | Change |
|------|--------|
| `scripts/check-coverage.ps1` | Fails when service or sqlite coverage drops below 55% |
| `internal/service/bookmark_test.go` | Extra create/list happy-path tests |

## Run tests & coverage

```powershell
cd learn/go/day-69

go test ./...

# Coverage profile (manual)
go test ./internal/service -coverprofile=service.out
go tool cover -func=service.out

# Automated gate
.\scripts\check-coverage.ps1
```

## Lint

```powershell
.\scripts\run-lint.ps1
```

## Local development

```powershell
Copy-Item .env.example .env
go run ./cmd/api
```

Next: **Day 70 capstone** combines lint + coverage + full integration suite.
