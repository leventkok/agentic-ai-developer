# Day 84 — Event-Driven Patterns

**Phase:** Caching & Messaging (Days 81–85)

Outbox relay, idempotent worker, and dead-letter queue. See [EVENT_PATTERNS.md](./EVENT_PATTERNS.md).

## What changed from Day 83

| Area | Change |
|------|--------|
| `migrations/007–008` | `outbox`, `processed_events` tables |
| `internal/messaging/outbox` | Durable event enqueue |
| `internal/messaging/relay` | Poll + publish loop in API |
| `internal/messaging/idempotency` | Dedup keys in worker |
| `internal/messaging/dlq` | Poison message quarantine |

## Run

```powershell
cd learn/go/day-84
go test ./...
go run ./cmd/api
go run ./cmd/worker   # with NATS_URL for separate process
```
