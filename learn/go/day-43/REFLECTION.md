# Retrospective — Day 40 (Context, Config & Middleware capstone)

## What went well
- Propagated `context.Context` from handlers into the store with cancellation and timeouts (Day 36)
- Typed config from env + `.env` for local dev, fail-fast validation at startup (Day 37)
- Middleware chain with request IDs, dependency injection, logging, and recovery (Day 38)
- Graceful shutdown with `http.Server.Shutdown` and read/write timeouts (Day 39)
- Request ID in every log line for traceability (Day 40)

## What was hard
- Middleware order — understanding which wrapper runs first on incoming requests
- Early `return nil` in `validate()` skipping later checks (Day 39)
- Naming consistency (`RequestIDFromContext` vs `RequestIdFromContext`)
- Knowing when to use `context.WithTimeout` in handler vs store vs shutdown

## What I'd change next time
- Add structured JSON logging instead of `fmt.Printf`
- Integration test that starts the real server and asserts `X-Request-ID`
- Move list timeout into config middleware or handler factory
- Persist bookmarks to SQLite/Postgres so restarts don't lose data

## Production patterns I'll reuse
- **Context as first param** on store/service methods
- **Config struct + Load() + validate()** — never scatter `os.Getenv` in handlers
- **Middleware chain** — cross-cutting concerns stay out of business logic
- **http.Server + Shutdown(ctx)** — always drain in-flight work on deploy
- **Request ID** — foundation for observability and support debugging
