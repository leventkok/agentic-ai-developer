# Day 83 — Domain Events & Worker

**Phase:** Caching & Messaging (Days 81–85)

Publish bookmark events from the service; consume them in `cmd/worker`.

## What changed from Day 82

| Area | Change |
|------|--------|
| `internal/messaging/nats` | Optional NATS broker |
| `internal/service/bookmark` | Publishes on create/update |
| `cmd/worker` | Background consumer with retries |

## Run

```powershell
cd learn/go/day-83
go test ./...

# API (in-memory bus — handlers run synchronously during publish)
go run ./cmd/api

# Worker with NATS
docker run -d --name nats -p 4222:4222 nats:2-alpine
$env:NATS_URL="nats://localhost:4222"
go run ./cmd/api
go run ./cmd/worker
```

Without `NATS_URL`, the memory bus is used inside the API process only.
