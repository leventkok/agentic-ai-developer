# Bookmarks API — 100-Day Go Capstone

**Version:** v1.0.0 · **Track complete**

Production-style Go service: HTTP + gRPC, auth, cache, messaging, observability, CI/CD.

## Quick start

```powershell
cd learn/go/day-100
copy .env.example .env
make test
make run
curl http://localhost:8080/health
curl http://localhost:8080/ready
```

## Final verification

```powershell
./scripts/verify-final-capstone.ps1
docker compose -f docker-compose.staging.yml up --build
```

## Portfolio

- [PORTFOLIO.md](./PORTFOLIO.md) — skills demonstrated
- [REFLECTION.md](./REFLECTION.md) — 100-day reflection
- [CAPSTONE.md](./CAPSTONE.md) — final checklist
