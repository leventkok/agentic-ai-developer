# Package layout — Day 56

## Responsibilities

| Package | Owns |
|---------|------|
| `domain` | Entities, domain errors, service inputs |
| `service` | Use cases (thin delegation to repos today) |
| `httpapi` | HTTP handlers, JSON DTOs, route registration |
| `ctxkey` | Request-scoped user (shared by httpapi + middleware) |
| `repository` | Persistence interfaces |
| `middleware` | Cross-cutting HTTP concerns |
| `cmd/api` | Composition root only |

## Checklist

- [x] domain package
- [x] service layer
- [x] httpapi migrated from handler/
- [x] thin main
- [x] tests green

## Note on import cycles

`middleware` cannot import `httpapi` (router imports middleware). User context lives in `internal/ctxkey` so both sides can share it.
