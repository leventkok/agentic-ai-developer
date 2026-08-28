# Day 49 — Connection Pooling

**Phase:** Databases (II) & Repositories (Days 46–50)

## Focus today
- Pool config via env: `DB_MAX_OPEN_CONNS`, `DB_MAX_IDLE_CONNS`, `DB_CONN_MAX_LIFETIME_MIN`
- `db.ConfigurePool` + `db.CollectPoolStats`
- Pool stats logged on graceful shutdown; DB closed after server stops

## Run
```powershell
cd learn/go/day-49
go mod tidy
go run ./cmd/api
# Ctrl+C — watch pool stats in shutdown log
```

See `.env.example` and `DATA_LAYER.md`.
