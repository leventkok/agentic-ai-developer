# Day 82 — Message Queues Overview

**Phase:** Caching & Messaging (Days 81–85)

Domain events and an in-process message bus. See [MESSAGING.md](./MESSAGING.md).

## What changed from Day 81

| Area | Package | Change |
|------|---------|--------|
| Events | `internal/messaging` | `Event`, `EventType`, dedup key |
| Bus | `internal/messaging/memory` | Publish / Subscribe for local dev |

## Run

```powershell
cd learn/go/day-82
go test ./...
```

Caching from Day 81 is unchanged (`internal/cache`, `internal/repository/cached`).
