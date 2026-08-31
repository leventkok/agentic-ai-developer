# Architecture — Bookmarks API

> Updated Day 92. See also `DOMAIN.md`, `DATA_LAYER.md`, `ASYNC_FLOW.md`, `PERFORMANCE.md`.

## System overview

```
Client → HTTP/gRPC → middleware → service → domain
                              ↓
                    repository (sqlite + cache + resilient)
                              ↓
                    outbox → relay → bus → worker
```

## Layers

| Layer | Package | Responsibility |
|-------|---------|----------------|
| Transport | `httpapi`, `grpcapi`, `middleware` | Protocol, auth, observability |
| Service | `service` | Use cases, event enqueue |
| Domain | `domain` | Rules, validation, entities |
| Data | `repository/*`, `cache`, `db` | Persistence, cache-aside |
| Async | `messaging/*`, `worker` | Outbox, idempotent consumers |

## Key decisions

1. **Clean architecture** — domain has zero framework imports (enforced by `internal/architecture` tests)
2. **Cache-aside** — Redis or in-memory; invalidate on writes
3. **Outbox pattern** — reliable event publish after DB commit
4. **Separate worker** — `cmd/worker` consumes NATS or in-memory bus

## Contracts

| Type | Location |
|------|----------|
| REST OpenAPI | [api/openapi.yaml](./api/openapi.yaml) |
| gRPC / Protobuf | [api/proto/bookmarks/v1/](./api/proto/bookmarks/v1/) |

## Dependency rule

Dependencies point inward only: transport → service → domain ← repository.
