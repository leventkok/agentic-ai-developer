# Day 76 — Dockerizing the Go API

**Phase:** Containers & CI/CD (Days 76–80)

Multi-stage Dockerfile that compiles `cmd/api` and runs it as PID 1.

## What changed from Day 75

| Area | File | Change |
|------|------|--------|
| Container | `Dockerfile` | Builder + minimal Alpine runtime |
| Ignore | `.dockerignore` | Smaller build context |

## Build & run

```powershell
cd learn/go/day-76
docker build -t bookmarks-api:day76 .
docker run --rm -p 8080:8080 -p 9090:9090 `
  -e JWT_SECRET=dev-docker-secret-at-least-32-chars `
  bookmarks-api:day76
```

Test:

```powershell
curl http://localhost:8080/bookmarks
curl http://localhost:8080/metrics
```

Migrations ship in the image at `/app/migrations`. SQLite uses `DB_PATH` (default `/tmp/bookmarks.db`).

## Tests (host)

```powershell
go test ./...
```
