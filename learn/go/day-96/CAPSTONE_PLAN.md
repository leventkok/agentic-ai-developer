# Capstone Plan — Day 96

## Scope

Extend the bookmarks MVP into a **production-ready service** demonstrating 100 days of Go engineering:

- HTTP + gRPC APIs with JWT auth
- SQLite persistence, Redis cache, outbox messaging
- Observability (logs, metrics, traces, pprof)
- Containers, CI/CD, team DX

## Requirements

| Area | Target |
|------|--------|
| **Availability** | `/health` liveness, `/ready` DB readiness |
| **Latency** | List p99 < 100ms on modest load (cached) |
| **Auth** | JWT on mutating routes; rate limit on auth |
| **Data** | Migrations 001–008 applied on startup |
| **Async** | Outbox relay + optional NATS worker |
| **Ops** | Prometheus `/metrics`, structured JSON logs |

## Architecture

See [ARCHITECTURE.md](./ARCHITECTURE.md), [ASYNC_FLOW.md](./ASYNC_FLOW.md).

## Milestones (Days 96–100)

| Day | Deliverable |
|-----|-------------|
| 96 | Plan + milestones (this doc) |
| 97 | Health/readiness endpoints + tests |
| 98 | Security review + smoke load |
| 99 | Staging deploy + CI artifacts |
| 100 | v1.0.0 release + portfolio |

## Out of scope

- Multi-region deployment
- Kubernetes manifests (post-100 learning path)
