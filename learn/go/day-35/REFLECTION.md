# MVP Reflection — Day 35

## What I learned
- How to split a REST API into layers: `cmd/` → `handler` → `store` → `model`
- CRUD handlers with correct status codes (201 create, 404 not found, 204 delete)
- Thread-safe in-memory storage with `sync.RWMutex` and a repository interface
- Centralized validation so handlers stay thin and errors stay consistent

## What was challenging
- Keeping packages in the right folders (`internal/model/` not `internal/bookmark.go`)
- `Update` with pointer fields (`*string`) vs create with plain strings
- Remembering that each `npm run dev` / new process resets in-memory data (Day 30 TS)
- Path aliases needing `tsc-alias` at build time — TypeScript doesn't rewrite them alone

## What I'd improve next
- Replace in-memory store with PostgreSQL or SQLite
- Add authentication (API keys or JWT)
- Pagination for GET /bookmarks (`?page=1&limit=20`)
- OpenAPI/Swagger docs for clients
- Integration tests that spin up the real HTTP server

## Technical debt (honest shortcuts)
- In-memory store — data lost on restart
- No request logging to file (only stdout middleware)
- No rate limiting or auth
- Validation is manual, not a library like `go-playground/validator`
- Recovery middleware returns plain text 500, not JSON problem responses
