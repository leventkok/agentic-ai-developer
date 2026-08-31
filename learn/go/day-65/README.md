# Day 65 — HTTP + gRPC Capstone

**Phase:** gRPC & Protocol Buffers (Days 61–65)

Runs bookmark operations over HTTP and gRPC from one process, sharing the same service layer.

## Local development

### 1. Configure environment

```powershell
cd learn/go/day-65
Copy-Item .env.example .env
```

Defaults: HTTP on `:8080`, gRPC on `:9090`.

### 2. Run both transports

```powershell
go run ./cmd/api
```

You should see:

- `HTTP on http://localhost:8080`
- `gRPC on localhost:9090`

### 3. Regenerate protobuf stubs (after `.proto` changes)

```powershell
.\scripts\generate-proto.ps1
```

Requires `protoc`, `protoc-gen-go`, and `protoc-gen-go-grpc` on your PATH.

### 4. Try HTTP auth + gRPC client

Terminal A — combined server:

```powershell
go run ./cmd/api
```

Terminal B — register/login over HTTP, then call gRPC:

```powershell
# register (first run)
curl -X POST http://localhost:8080/auth/register -H "Content-Type: application/json" -d "{\"email\":\"demo@example.com\",\"password\":\"password123\"}"

# demo client: lists bookmarks, logs in via HTTP, creates via gRPC
go run ./cmd/grpcclient -http http://localhost:8080 -addr localhost:9090
```

Or run gRPC only:

```powershell
go run ./cmd/grpcserver
```

### 5. Tests

```powershell
go test ./...
```

## Architecture

| Layer | Package | Role |
|-------|---------|------|
| Transport | `internal/httpapi` | REST JSON handlers |
| Transport | `internal/grpcapi` | gRPC handlers + interceptors |
| Application | `internal/service` | Shared business orchestration |
| Domain | `internal/domain` | Rules and validation |
| Wiring | `internal/app/wire.go` | Composition root for HTTP + gRPC |

Protected gRPC RPCs (`CreateBookmark`, `DeleteBookmark`) require `authorization: Bearer <token>` metadata. The server interceptor validates JWTs and attaches the user to context before handlers run.

## Standalone commands

| Command | Purpose |
|---------|---------|
| `cmd/api` | HTTP + gRPC together (capstone) |
| `cmd/grpcserver` | gRPC only |
| `cmd/grpcclient` | Demo client with request-id interceptor |

See [CAPSTONE.md](./CAPSTONE.md) for the phase checklist.
