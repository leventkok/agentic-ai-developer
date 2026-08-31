# Day 71 — Structured Logging

**Phase:** Observability & Resilience (Days 71–75)

JSON structured logs with `slog`, request correlation fields, and sensible log levels.

## What changed from Day 70

| Area | Package | Change |
|------|---------|--------|
| Logger | `internal/observability/log` | JSON `slog` handler per environment |
| Middleware | `internal/middleware/logging.go` | Logs `request_id`, method, path, status, duration |
| Levels | middleware | INFO/WARN/ERROR by HTTP status |

## Run

```powershell
cd learn/go/day-71
go test ./...
go run ./cmd/api
# Example log line (JSON):
# {"time":"...","level":"INFO","msg":"http request","request_id":"...","method":"GET","path":"/bookmarks","status":200,"duration":...}
```

Passwords and tokens are never logged. Auth handlers log errors without credentials.

## Quality gate (from Day 70)

```powershell
.\scripts\verify-quality.ps1
```
