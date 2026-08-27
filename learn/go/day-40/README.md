# Day 40 — Production-Minded Bookmarks API (Capstone)

**Phase:** Context, Config & Middleware (Days 36–40)

## Run
```powershell
cd learn/go/day-40
copy .env.example .env   # first time only
go run ./cmd/api
```

## Architecture
```
cmd/api/main.go           → config load, routes, http.Server, graceful shutdown
internal/config/          → typed env config + .env loader
internal/middleware/      → Recovery → Logging → RequestID → InjectDeps
internal/handler/         → HTTP layer, passes r.Context() to store
internal/store/           → context-aware repository + slow List for timeout demos
```

## Environment variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP listen port |
| `ENV` | `development` | `development`, `staging`, or `production` |
| `LIST_TIMEOUT_MS` | `100` | Handler timeout for slow `GET /bookmarks` |
| `READ_TIMEOUT_SEC` | `5` | Max time to read request |
| `WRITE_TIMEOUT_SEC` | `10` | Max time to write response |
| `SHUTDOWN_TIMEOUT_SEC` | `15` | Max wait for in-flight requests on shutdown |

## Endpoints

| Method | Route | Status |
|--------|-------|--------|
| GET | /bookmarks | 200 / 408 |
| POST | /bookmarks | 201 / 400 |
| GET | /bookmarks/{id} | 200 / 404 |
| PATCH | /bookmarks/{id} | 200 / 404 |
| DELETE | /bookmarks/{id} | 204 / 404 |

## Tests
```powershell
go test ./...
go vet ./...
```

## Manual verify
```powershell
curl.exe -i http://localhost:8080/bookmarks
# Expect: X-Request-ID header + log line like [abc123] GET /bookmarks 200 12ms

# Shutdown: Ctrl+C while server running
# Expect: shutdown signal received → server stopped cleanly
```
