# Reflection — 100-Day Go Journey

## What I built

A production-style bookmarks API spanning HTTP and gRPC, with SQLite persistence, JWT auth, Redis caching, outbox messaging, full observability stack, Docker deployment, and CI/CD gates.

## Biggest growth areas

1. **Clean architecture** — keeping domain pure and testing rules without a database
2. **Operability** — metrics, structured logs, and health probes as first-class features
3. **Async patterns** — outbox + idempotent consumers instead of fire-and-forget publishes
4. **Team workflow** — PR templates, changelogs, and Makefile-driven DX

## What was hardest

- Wiring dependency injection without framework magic (`app.Wire`)
- Understanding when cache invalidation races are acceptable vs when outbox is required
- Benchmark-driven optimization vs premature micro-optimizations

## What I'd do differently on a greenfield project

- Start with OpenAPI contract and architecture tests on day one
- Use PostgreSQL from the beginning if production is the goal
- Add health/readiness endpoints before the first deploy

## Next steps

- Deploy to a cloud environment with managed Redis and NATS
- Learn Kubernetes for orchestration
- Deepen gRPC streaming and protobuf evolution practices

## Interview talking points

- End-to-end feature: create bookmark → outbox → worker → idempotent handler
- Profile-guided optimization with pprof and `-benchmem`
- Security: JWT, rate limiting, parameterized SQL, production config validation
