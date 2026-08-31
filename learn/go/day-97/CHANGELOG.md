# Changelog

All notable changes to the bookmarks API are documented here.

## [0.2.0] — 2026-08-31

### Added
- Redis cache-aside and NATS/outbox messaging
- pprof debug port and benchmark regression guard
- Makefile, OpenAPI contract, PR template, release scripts
- Idempotent worker with DLQ

### Deprecated
- `POST /bookmarks/bulk` — use single creates; removal in v1.0

## [0.1.0] — 2026-08-31

### Added
- HTTP + gRPC bookmark CRUD with JWT auth
- SQLite persistence, Docker, CI, observability stack

[0.2.0]: ./RELEASE_NOTES/v0.2.0.md
[0.1.0]: https://example.com/releases/tag/v0.1.0
