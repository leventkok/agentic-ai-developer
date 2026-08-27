# Day 38 — Middleware Chaining

**Phase:** Context, Config & Middleware (Days 36–40)

## Run
```powershell
cd learn/go/day-38
go run ./cmd/api
```

---

## Tasks — which file to edit

| Task | What you do | File to edit |
|------|-------------|--------------|
| **1** | Add request ID to each request + context | `internal/middleware/requestid.go` **(NEW)** |
| **2** | Inject config/store into request context | `internal/middleware/deps.go` **(NEW)** |
| **3** | Order middleware in one helper | `internal/middleware/chain.go` |
| **4** | Test request ID header + middleware | `internal/middleware/requestid_test.go` **(NEW)** |
| **5** | Wire middleware chain in main | `cmd/api/main.go` |

**Do not edit today:** `handler/bookmarks.go`, `store/*`, `config/*` (unless a test breaks)

---

## Task 1 — Request ID middleware

**File:** `internal/middleware/requestid.go`

**Logic:** Every request gets a unique ID. Put it in context so handlers/logging can use it later.

```go
package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid" // OR use crypto/rand — see lesson
)

type ctxKey string

const RequestIDKey ctxKey = "requestID"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.NewString()
		w.Header().Set("X-Request-ID", id)
		ctx := context.WithValue(r.Context(), RequestIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
```

> Simpler option without uuid package: generate 8 random hex chars with `crypto/rand`.

---

## Task 2 — Dependency injection middleware

**File:** `internal/middleware/deps.go`

**Logic:** Instead of handlers reaching for globals, middleware attaches shared deps to context.

```go
package middleware

import (
	"context"
	"net/http"

	"learn/go/day-38/internal/config"
	"learn/go/day-38/internal/store"
)

type depsKey struct{}

type Deps struct {
	Config config.Config
	Store  store.BookmarkRepository
}

func InjectDeps(cfg config.Config, s store.BookmarkRepository) func(http.Handler) http.Handler {
	deps := Deps{Config: cfg, Store: s}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), depsKey{}, deps)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func DepsFromContext(ctx context.Context) (Deps, bool) {
	deps, ok := ctx.Value(depsKey{}).(Deps)
	return deps, ok
}
```

---

## Task 3 — Compose the chain

**File:** `internal/middleware/chain.go`

**Logic:** Order matters — outer middleware runs first on the way in:

```
Request → Recovery → Logging → RequestID → InjectDeps → Handler
```

Add a named helper so `main` stays readable:

```go
func DefaultStack(cfg config.Config, store store.BookmarkRepository, mux http.Handler) http.Handler {
	return Chain(
		mux,
		InjectDeps(cfg, store),
		RequestID,
		Logging,
		Recovery,
	)
}
```

**Why this order?**
- **Recovery** outermost — catches panics from everything inside
- **Logging** — logs status after handler runs
- **RequestID** — every log line can include the ID
- **InjectDeps** — closest to handler, deps ready when handler runs

---

## Task 4 — Test middleware

**File:** `internal/middleware/requestid_test.go`

**Logic:** Test middleware in isolation with `httptest` — no real server needed.

```go
func TestRequestID_SetsHeader(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, ok := r.Context().Value(RequestIDKey).(string)
		if !ok || id == "" {
			t.Fatal("request ID missing from context")
		}
	})
	h := RequestID(next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Header().Get("X-Request-ID") == "" {
		t.Fatal("X-Request-ID header not set")
	}
}
```

---

## Task 5 — Wire in main

**File:** `cmd/api/main.go`

Replace manual `Chain(...)` with:

```go
root := middleware.DefaultStack(cfg, s, mux)
```

Remove duplicate store/handler wiring if you move to context-based deps later (optional for today).

---

## Verify
```powershell
go test ./internal/middleware/... -v
go test ./...
go run ./cmd/api
curl.exe -i http://localhost:8080/bookmarks
# Look for X-Request-ID header
```
