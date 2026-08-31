# Day 99 — Deployment and CI/CD

Staging stack with health checks and CI artifact upload.

```powershell
cd learn/go/day-99
docker compose -f docker-compose.staging.yml up --build
./scripts/verify-deploy.ps1
```

See [DEPLOY.md](./DEPLOY.md).
