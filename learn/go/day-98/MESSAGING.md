# Message Queues Overview — Day 82

## Concepts

| Term | Meaning |
|------|---------|
| **Producer** | Publishes events after domain changes |
| **Consumer / Worker** | Processes events off the HTTP path |
| **Topic / Subject** | Named channel (`bookmark.created`) |
| **Ack** | Consumer confirms successful processing |
| **At-least-once** | Duplicates possible — handlers must be idempotent |

## Async use cases in this app

- Search indexing after bookmark create
- Webhook notifications
- Audit enrichment (complementing sync audit table)

## Package layout

```
internal/messaging/
  event.go      — domain events + deduplication key
  bus.go        — Bus interface
  memory/       — in-process bus for tests and local dev
```

Brokers (NATS, RabbitMQ, Kafka) differ in durability and ops cost. We start with an in-memory bus, then add NATS in Day 83.

## Idempotent consumers (preview)

Every event carries an `id` used as a deduplication key. Day 84 stores processed keys in SQLite so redeliveries are safe.
