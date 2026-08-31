# Containers & CI/CD Capstone — Day 80

**Phase:** Containers & CI/CD (Days 76–80)

Deployable container image + CI pipeline from laptop to GitHub Actions.

## Phase recap

| Day | Focus | Deliverable |
|-----|-------|-------------|
| 76 | Dockerize | Multi-stage `Dockerfile` |
| 77 | Optimize | Layer cache, `-ldflags`, non-root user |
| 78 | GHA basics | Test workflow + module cache |
| 79 | Full CI | lint, vet, build, integration, coverage artifact |
| 80 | **Capstone** | `docker-compose`, `DEPLOY.md`, docker CI job |

## Capstone checklist

- [x] Multi-stage Docker build for `cmd/api`
- [x] `.dockerignore` for lean context
- [x] `docker compose` local stack
- [x] GitHub Actions: test, lint, docker build, integration
- [x] `DEPLOY.md` operator docs
- [x] `VERSION` = v0.1.0
- [x] `scripts/verify-deploy.ps1`

## Verify

```powershell
cd learn/go/day-82
go test ./...
.\scripts\verify-deploy.ps1
```

## What you built

A path from **code → CI → container → runbook** — the same flow teams use before Kubernetes or cloud deploy.
