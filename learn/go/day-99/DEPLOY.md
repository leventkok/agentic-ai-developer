# Deploy — Bookmarks API v1.0.0

## Staging (production-like)

```powershell
cd learn/go/day-99
docker compose -f docker-compose.staging.yml up --build -d
curl http://localhost:8080/health
curl http://localhost:8080/ready
curl http://localhost:8080/metrics
docker compose -f docker-compose.staging.yml down
```

## Rollback

1. Note previous image tag: `bookmarks-api:v0.2.0`
2. Stop current stack: `docker compose down`
3. Deploy previous tag with same env vars
4. Verify: `curl /health` and `curl /ready`

## CI artifacts

GitHub Actions uploads `bookmarks-api` binary and coverage profile on main branch.

## Environment

| Variable | Staging | Production |
|----------|---------|------------|
| `ENV` | staging | production |
| `JWT_SECRET` | staging secret (32+ chars) | unique secret |
| `PPROF_PORT` | 6060 | empty (disabled) |
| `REDIS_URL` | redis://redis:6379/0 | managed Redis |
| `NATS_URL` | nats://nats:4222 | managed NATS |

## Migrations

Applied automatically on API startup. For rollback, restore DB snapshot — migrations are forward-only.

## Verify locally

```powershell
./scripts/verify-deploy.ps1
./scripts/verify-final-capstone.ps1   # Day 100+
```
