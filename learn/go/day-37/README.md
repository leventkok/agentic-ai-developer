# Day 37 — Environment Config

**Phase:** Context, Config & Middleware (Days 36–40)  
**Base:** Day 36 Bookmarks API with context support

## Run
```powershell
cd learn/go/day-37
go run ./cmd/api
```

## Today's goal
Load **typed configuration** from environment variables, support local `.env`, and **fail fast** on bad config at startup.

## Folder layout (new today)
```
internal/config/
  config.go       ← Config struct + Load() + validation
  config_test.go  ← tests for defaults and fail-fast
.env.example      ← committed template (no secrets)
.env              ← local only (gitignored)
```

## Steps

### Step 1 — Typed `Config` struct
File: `internal/config/config.go`

### Step 2 — Read env vars with defaults
Use `os.Getenv` + fallbacks.

### Step 3 — Load `.env` locally (optional file)
Simple line parser: `KEY=VALUE`, skip `#` comments.

### Step 4 — Validate on startup
Return error if config is invalid; `main` calls `log.Fatal`.

### Step 5 — Wire `main.go`
Replace hardcoded `:8080` and `100*time.Millisecond` with config values.

### Step 6 — Tests
Test defaults, invalid values, and `.env` loading.

See lesson in chat for full code and logic explanations.
