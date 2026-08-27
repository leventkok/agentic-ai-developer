# Day 36 — context.Context Fundamentals

**Phase:** Context, Config & Middleware (Days 36–40)  
**Base:** Bookmarks API from Day 35 (compiles as-is — you will add context today)

## Run (baseline)
```powershell
cd learn/go/day-36
go run ./cmd/api
```

## Today's goal
Propagate `context.Context` from HTTP handlers into the store, honor cancellation, and bound slow work with timeouts.

## Folder layout
```
cmd/api/              → entry point
internal/handler/     → pass r.Context() to store
internal/store/       → ctx as first param + ctx.Done() in loops
internal/model/       ← unchanged today
internal/validation/  ← unchanged today
internal/middleware/    ← unchanged today
```

## Steps (do in order)

### Step 1 — Update the repository interface
File: `internal/store/repository.go`

Add `context.Context` as the **first parameter** on every method. Return `error` where cancellation or not-found can happen:

```go
List(ctx context.Context) ([]model.Bookmark, error)
Get(ctx context.Context, id int) (model.Bookmark, error)
Create(ctx context.Context, req model.CreateBookmarkRequest) (model.Bookmark, error)
Update(ctx context.Context, id int, req model.UpdateBookmarkRequest) (model.Bookmark, error)
Delete(ctx context.Context, id int) error
```

Add sentinel errors in the same file:
```go
var ErrNotFound = errors.New("bookmark not found")
```

### Step 2 — Honor ctx in MemoryStore
File: `internal/store/memory.go`

- Update every method signature to match the interface.
- In `List`, simulate slow work: `time.Sleep(10 * time.Millisecond)` **per bookmark** inside the loop.
- Before each sleep (and before expensive steps), check cancellation:

```go
select {
case <-ctx.Done():
    return nil, ctx.Err()
default:
}
```

- `Get` / `Update` / `Delete`: return `ErrNotFound` when id missing (not `(zero, false)`).
- Fix `memory_test.go` to pass `context.Background()`.

### Step 3 — Thread context from handlers
File: `internal/handler/bookmarks.go`

Pass `r.Context()` into every store call:

```go
list, err := h.Store.List(r.Context())
```

Handle errors:
- `errors.Is(err, context.Canceled)` or `context.DeadlineExceeded` → **408 Request Timeout** (or 499)
- `errors.Is(err, store.ErrNotFound)` → **404**
- other errors → **500**

Optional: wrap slow `List` with a handler timeout:
```go
ctx, cancel := context.WithTimeout(r.Context(), 100*time.Millisecond)
defer cancel()
list, err := h.Store.List(ctx)
```

### Step 4 — Write context tests
File: `internal/store/context_test.go`

**Test A — WithCancel:** start `List` in a goroutine, cancel ctx after 1 item, assert `context.Canceled`.

**Test B — WithTimeout:** seed store with many bookmarks, call `List` with `context.WithTimeout(..., 5*time.Millisecond)`, assert `context.DeadlineExceeded`.

## Verify
```powershell
go test ./...
go vet ./...
go run ./cmd/api
```

Manual: seed 20+ bookmarks, hit `GET /bookmarks` — should still work. Then add a short timeout in the handler and confirm 408 when List is too slow.

## Dictionary
| Term | Meaning |
|------|---------|
| `context.Context` | Deadlines, cancellation, request-scoped values |
| `ctx.Done()` | Channel closed when canceled or timed out |
| `WithCancel` | Manual cancel via `cancel()` function |
| `WithTimeout` | Auto-cancel after duration |
| Goroutine leak | Work continues after nobody needs the result |
