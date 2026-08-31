# Security practices — Day 54

## Rate limiting

Auth routes are limited to `AUTH_RATE_LIMIT_PER_MIN` requests per IP per minute.

Returns **429 Too Many Requests** when exceeded. In-memory store resets on process restart — use Redis for multi-instance production.

## HTTPS

Development uses plain HTTP. In production:

- Terminate TLS at a reverse proxy (nginx, Caddy, cloud load balancer), or
- Use `http.Server` with TLS certificates

Never send JWTs or passwords over plaintext in production.

## Dependency audit

```powershell
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```

Run in CI weekly or on every release. Update modules when vulnerabilities are reported.

## Input validation

Authentication proves *who* the caller is, not that their data is safe. Keep validation on register, login, and bookmark payloads.

## Least privilege

Default role is `member`. Grant `admin` only when needed.
