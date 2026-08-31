# Event-Driven Patterns — Day 84

## Outbox pattern

1. HTTP handler commits bookmark to SQLite **and** inserts row into `outbox`.
2. `relay` goroutine polls unpublished rows and publishes to the bus.
3. On success, row is marked `published_at`.

This avoids losing events when the DB commit succeeds but publish fails.

## Idempotent consumers

Worker wraps handlers with `idempotency.Wrap`: `processed_events.dedup_key` ensures duplicate deliveries are no-ops.

## Dead letter queue

After retry exhaustion, failed events land in an in-memory DLQ (`internal/messaging/dlq`) for inspection.

## Saga (sketch)

A multi-step bookmark import could emit:

`import.started` → `bookmark.created` (×N) → `import.completed`

Each step is a local transaction plus outbox row; compensating events roll back prior steps on failure.
