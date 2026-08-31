# Bookmarks API — Day 95 Team Practices Capstone

Production-style Go service with cache, messaging, observability, and team DX.

## Quick start

```powershell
cd learn/go/day-95
copy .env.example .env
make test
make run
```

| Command | Purpose |
|---------|---------|
| `make test` | Unit tests |
| `make lint` | golangci-lint |
| `make run` | HTTP + gRPC API |
| `make verify` | Quality gate |
| `make hooks` | Install pre-commit hooks |

## Docs

- [CONTRIBUTING.md](./CONTRIBUTING.md) — bootstrap & workflow
- [CODE_REVIEW.md](./CODE_REVIEW.md) — PR checklist
- [ARCHITECTURE.md](./ARCHITECTURE.md) — system design
- [CHANGELOG.md](./CHANGELOG.md) — version history
- [CAPSTONE.md](./CAPSTONE.md) — Day 95 checklist

## Release

```powershell
./scripts/release.ps1
./scripts/verify-team-practices.ps1
```

Current version: **v0.2.0** (see [VERSION](./VERSION))
