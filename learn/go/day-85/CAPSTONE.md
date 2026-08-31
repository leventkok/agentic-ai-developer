# Day 85 — Caching & Messaging Capstone

**Phase complete:** Days 81–85

This day ties together cache-aside, domain events, outbox relay, idempotent workers, and optional Redis/NATS infrastructure.

## Checklist

- [x] Cache-aside for List/Get (`internal/repository/cached`)
- [x] Domain events on create/update
- [x] Outbox + relay (`migrations/007`, `internal/messaging/relay`)
- [x] Idempotent worker + DLQ (`migrations/008`, `cmd/worker`)
- [x] Async flow diagram — [ASYNC_FLOW.md](./ASYNC_FLOW.md)

## Verify

```powershell
cd learn/go/day-85
./scripts/verify-cache-messaging.ps1
```

## Docker stack (optional)

```powershell
docker compose up --build
# API:     http://localhost:8080
# Worker:  consumes from NATS
# Redis:   localhost:6379
# NATS:    localhost:4222
```

## Next phase

Days 86–90: Performance & profiling (pprof, benchmarks, load testing).
