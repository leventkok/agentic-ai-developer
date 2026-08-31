# Security Review — Day 98

Checklist completed for bookmarks capstone.

| Check | Status | Notes |
|-------|--------|-------|
| JWT on mutating routes | ✓ | `RequireAuth` middleware |
| Rate limit on auth | ✓ | `AUTH_RATE_LIMIT_PER_MIN` |
| Production JWT secret | ✓ | Config validation blocks default in prod |
| SQL injection | ✓ | Parameterized queries only |
| Input validation | ✓ | Domain + validation packages |
| Secrets in repo | ✓ | `.env` gitignored |
| Dependency scan | ○ | Run `go mod tidy`; CI builds clean |

## Resilience

| Check | Status |
|-------|--------|
| HTTP read/write timeouts | ✓ |
| Circuit breaker on repo reads | ✓ | `resilient` wrapper |
| Retry with backoff | ✓ | `internal/resilience` |
| Graceful shutdown | ✓ | `cmd/api/main.go` |

## Observability

| Check | Status |
|-------|--------|
| Structured JSON logs | ✓ |
| Prometheus `/metrics` | ✓ |
| OpenTelemetry tracing | ✓ |
| pprof debug port | ✓ | Disable in prod (`PPROF_PORT=`) |

## Smoke load

```powershell
./scripts/load-test.ps1 -Requests 200
./scripts/verify-observability.ps1
```
