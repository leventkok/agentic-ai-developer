# Async flow — Day 85 capstone

## Create bookmark (happy path)

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant Cache
    participant DB
    participant Outbox
    participant Relay
    participant Bus
    participant Worker

    Client->>API: POST /bookmarks
    API->>DB: INSERT bookmark
    API->>Outbox: INSERT event row
    API->>Cache: DELETE list:* keys
    API-->>Client: 201 Created

    loop every 500ms
        Relay->>Outbox: SELECT unpublished
        Relay->>Bus: Publish bookmark.created
        Relay->>Outbox: SET published_at
    end

    Bus->>Worker: deliver event
    Worker->>DB: INSERT processed_events (dedup)
    Worker->>Worker: handle (log/index/webhook)
```

## Cache read path

```mermaid
flowchart LR
    GET["GET /bookmarks"] --> Cache{cache hit?}
    Cache -->|yes| Response
    Cache -->|no| DB[(SQLite)]
    DB --> Populate[SET cache TTL]
    Populate --> Response
```

## Failure modes

| Scenario | Mitigation |
|----------|------------|
| Publish fails after DB commit | Outbox relay retries |
| Duplicate delivery | `processed_events` dedup key |
| Handler always fails | DLQ after max retries |
| Stale cache | TTL + invalidation on writes |

## Race window

Between DB update and cache invalidation, a reader may see old data for at most `CACHE_TTL_SEC`. Writes always invalidate eagerly.
