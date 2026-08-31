# Day 77 — Optimized Multi-Stage Docker Builds

**Phase:** Containers & CI/CD (Days 76–80)

Cached module layers, static binary with stripped symbols, non-root runtime user.

## What changed from Day 76

| Area | Change |
|------|--------|
| Dockerfile | `-ldflags="-s -w"`, `-trimpath`, `CGO_ENABLED=0` |
| Cache | `go.mod` / `go.sum` copied before source |
| Security | `USER 65532:65532` |
| Scripts | `build-image.ps1`, optional `scan-image.ps1` (Trivy) |

## Build

```powershell
cd learn/go/day-77
.\scripts\build-image.ps1
.\scripts\scan-image.ps1
```

## Run

```powershell
docker run --rm -p 8080:8080 -p 9090:9090 `
  -e JWT_SECRET=dev-docker-secret-at-least-32-chars `
  bookmarks-api:day77
```
