# Day 89 — Database and HTTP Performance

**Phase:** Performance (Days 86–90)

Query plan inspection and keep-alive HTTP client.

## Changes

| Area | Package |
|------|---------|
| EXPLAIN QUERY PLAN | `internal/db/explain.go` |
| Keep-alive client | `internal/httpclient` |

```powershell
cd learn/go/day-89
go test ./internal/db/... -run TestExplain
go test ./...
```
