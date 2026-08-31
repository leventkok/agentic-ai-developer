# Day 81 — Cache-Aside with Redis or In-Memory

**Phase:** Caching & Messaging (Days 81–85)

Cache-aside for bookmark List/Get with TTL and invalidation on writes.

## What changed from Day 80

| Area | Package | Change |
|------|---------|--------|
| Cache | `internal/cache` | Store interface + memory + Redis |
| Repository | `internal/repository/cached` | Cache-aside wrapper |
| Config | `REDIS_URL`, `CACHE_TTL_SEC` | External or in-memory cache |

## Run

```powershell
cd learn/go/day-81
go test ./...
go run ./cmd/api
```

With Redis:

```powershell
$env:REDIS_URL="redis://localhost:6379/0"
go run ./cmd/api
```

## Pattern

1. **Read:** cache → on miss DB → populate cache  
2. **Write:** update DB → delete cache keys  
3. **TTL:** limits staleness (`CACHE_TTL_SEC`, default 60)
