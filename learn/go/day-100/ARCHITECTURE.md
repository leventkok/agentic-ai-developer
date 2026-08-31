# Architecture — Bookmarks API

> Day 95 capstone. Full stack from Days 56–90 plus team practices.

## Overview

```
Client → HTTP/gRPC → middleware → service → domain
                              ↓
              repository (sqlite + cache + resilient)
                              ↓
                    outbox → relay → bus → worker
```

## Layers

| Layer | Package |
|-------|---------|
| Transport | `httpapi`, `grpcapi`, `middleware` |
| Service | `service` |
| Domain | `domain` |
| Data | `repository/*`, `cache` |
| Async | `messaging/*`, `worker` |

## Contracts

- REST: [api/openapi.yaml](./api/openapi.yaml)
- gRPC: [api/proto/bookmarks/v1/](./api/proto/bookmarks/v1/)
