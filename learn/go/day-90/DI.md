# Dependency Injection — Day 58

## Why explicit DI?

Constructors state what a type needs. Readers see coupling immediately; tests swap fakes without databases.

```
cmd/api/main.go          ← thin entrypoint
    └── internal/app/wire.go   ← composition root (only concrete wiring)
            ├── sqlite.Store
            ├── sqlite.AuthStore
            ├── service.BookmarkService(repo, timeout)
            ├── service.AuthService(authRepo)
            └── httpapi.NewRouter(...)
```

## Interface boundaries

| Dependency | Interface | Production | Tests |
|------------|-----------|------------|-------|
| Bookmarks | `repository.Bookmarks` | `sqlite.Store` | `fake.Bookmarks` |
| Auth | `repository.Auth` | `sqlite.AuthStore` | fake (TODO) |
| Time | `clock.Clock` | `clock.RealClock` | fixed `fakeClock{now: ...}` |
| JWT | `repository.Auth` (token side) | `auth.TokenService` via store | stub |

## Anti-pattern to remove

**Service locator** — hiding dependencies in context/globals:

```go
// middleware/deps.go — InjectDeps puts cfg+repo in request context
// Nothing in handlers uses DepsFromContext today, but the pattern hides deps.
// Day 58: remove InjectDeps from DefaultStack once all handlers get explicit deps.
```

Prefer passing `*service.BookmarkService` into handlers (already done) over context lookup.

## Today's refactor targets

- [ ] **`internal/app/wire.go`** — single composition root; `main` calls `app.Wire(cfg)`
- [ ] **`internal/clock/clock.go`** — inject `Clock` into memory store / JWT (replace `time.Now()`)
- [ ] **`internal/service/testing/fake/`** — in-memory fakes implementing repository interfaces
- [ ] **`internal/service/*_test.go`** — service tests with fakes (no HTTP, no SQLite)
- [ ] **Remove `middleware.InjectDeps`** from `DefaultStack` if unused
- [ ] **Constructor audit** — every `New*` lists dependencies as parameters, creates nothing global

## Verify

```powershell
cd learn/go/day-59
go test ./internal/service/...   # fakes, no integration DB
go test ./...
go build ./cmd/api
```

## Test pyramid after Day 58

```
        /  HTTP (httpapi)  \        few, slow
       /  Service + fakes  \       many, fast
      /  Domain rules       \      pure, fastest
```
