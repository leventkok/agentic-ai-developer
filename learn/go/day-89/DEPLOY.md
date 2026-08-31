# Deploy — Bookmarks API (Day 80)

## Build image

```powershell
cd learn/go/day-89
docker build -t bookmarks-api:v0.1.0 .
```

Or:

```powershell
docker compose build
```

## Run container

```powershell
docker run --rm -p 8080:8080 -p 9090:9090 \
  -e JWT_SECRET=your-production-secret-min-32-chars \
  -e ENV=production \
  bookmarks-api:v0.1.0
```

## Docker Compose (local)

```powershell
docker compose up --build
curl http://localhost:8080/bookmarks
curl http://localhost:8080/metrics
docker compose down
```

## Required environment

| Variable | Default | Notes |
|----------|---------|-------|
| `PORT` | 8080 | HTTP listen port |
| `GRPC_PORT` | 9090 | gRPC listen port |
| `JWT_SECRET` | — | Required, min 32 chars in production |
| `DB_PATH` | `/tmp/bookmarks.db` | SQLite file inside container |
| `ENV` | development | `production` enforces JWT change |

## Release tag

Version file: `VERSION` → `v0.1.0`

```powershell
git tag learn/go/day-89/v0.1.0
```

## CI

Copy `.github/workflows/ci.yml` to repository root to run test, lint, docker build, and integration on push.

Local verify:

```powershell
.\scripts\verify-deploy.ps1
```
