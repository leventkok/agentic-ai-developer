# Day 39 — Graceful Shutdown

**Phase:** Context, Config & Middleware (Days 36–40)

## Run
```powershell
cd learn/go/day-39
go run ./cmd/api
# Ctrl+C → should drain and exit cleanly
```

---

## Tasks — which file to edit

| Task | What you do | File to edit |
|------|-------------|--------------|
| **1** | Add shutdown timeout config | `internal/config/config.go` |
| **2** | Use `http.Server` with read/write timeouts | `cmd/api/main.go` |
| **3** | Trap SIGINT/SIGTERM | `cmd/api/main.go` |
| **4** | Call `server.Shutdown(ctx)` to drain requests | `cmd/api/main.go` |

**Do not edit today:** `handler/*`, `store/*`, `middleware/*` (unless tests break)

---

See chat lesson for full code + logic per task.
