# Day 79 — Full Go CI Pipeline

**Phase:** Containers & CI/CD (Days 76–80)

Test, vet, build, lint, integration, and coverage artifacts in GitHub Actions.

## What changed from Day 78

| Job | Steps |
|-----|-------|
| `test` | `go test`, `go vet`, `go build`, coverage upload |
| `lint` | golangci-lint-action |
| `integration` | `-tags=integration` httpapi tests |

## Local parity

```powershell
cd learn/go/day-79
go test ./...
go vet ./...
go build ./...
.\scripts\verify-quality.ps1
go test -tags=integration ./internal/httpapi/...
```

Copy workflow to repo root `.github/workflows/` to activate on GitHub.
