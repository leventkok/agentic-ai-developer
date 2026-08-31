# Team Practices Capstone — Day 95

## Checklist

- [x] PR template and [CODE_REVIEW.md](./CODE_REVIEW.md)
- [x] [CONTRIBUTING.md](./CONTRIBUTING.md) bootstrap guide
- [x] godoc, [CHANGELOG.md](./CHANGELOG.md), [api/openapi.yaml](./api/openapi.yaml)
- [x] [Makefile](./Makefile), optional hooks and direnv
- [x] Release [v0.2.0](./VERSION) with [RELEASE_NOTES](./RELEASE_NOTES/v0.2.0.md)
- [x] Self-review checklist below

## Self-review (mock PR)

| Item | Status |
|------|--------|
| Tests pass (`go test ./...`) | ✓ |
| Quality gate (`verify-quality.ps1`) | ✓ |
| Auth on protected routes | ✓ |
| Domain free of HTTP/SQL imports | ✓ |
| OpenAPI matches router paths | ✓ |
| Deprecation headers on bulk endpoint | ✓ |

## Post-100 gaps

- Production Redis/NATS runbooks in staging
- OpenAPI response schemas for all status codes
- Automated release tagging in CI

## Verify

```powershell
./scripts/verify-team-practices.ps1
```

## Next phase

Days 96–100: Final capstone and course completion.
