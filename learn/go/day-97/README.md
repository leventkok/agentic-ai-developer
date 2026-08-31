# Day 97 — Building Core Features

**Phase:** Final Capstone (Days 96–100)

Kubernetes-style **liveness** (`GET /health`) and **readiness** (`GET /ready`) probes with DB ping.

```powershell
cd learn/go/day-97
go test ./internal/httpapi/... -run Health
go test ./...
curl http://localhost:8080/health
curl http://localhost:8080/ready
```
