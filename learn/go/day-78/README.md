# Day 78 — GitHub Actions Basics

**Phase:** Containers & CI/CD (Days 76–80)

CI workflow template running Go tests on push/PR with module cache.

## What changed from Day 77

| Area | Change |
|------|--------|
| CI | `.github/workflows/ci.yml` with `actions/setup-go` + test job |

## Local equivalent

```powershell
cd learn/go/day-78
go test ./...
```

## Enable in GitHub

Copy `.github/workflows/ci.yml` to the **repository root** `.github/workflows/` (GitHub only reads root workflows).

Path filters target `learn/go/day-78/**` so other tracks are unaffected.
