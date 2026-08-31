# Day 54 — Security Best Practices

**Phase:** Auth & Security (Days 51–55)

Builds on Day 53 RBAC. Adds **rate limiting** and security hygiene.

## What you built

- `internal/middleware/ratelimit.go` — per-IP limit on auth routes (429)
- `AUTH_RATE_LIMIT_PER_MIN` in config
- Input validation unchanged but documented as required even after auth
- Dependency audit guidance with `govulncheck`

## Run

```powershell
cd learn/go/day-54
go mod tidy
go run ./cmd/api
go test ./...

# Dependency audit
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```

## Security checklist

- Validate all input (auth does not imply trust)
- Use HTTPS in production (TLS termination)
- Rate limit `/auth/register` and `/auth/login`
- Run `govulncheck` regularly in CI

See `SECURITY.md` for details.

**Next:** Day 55 capstone — full auth flow tests + threat review.
