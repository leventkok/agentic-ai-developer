# Day 64 — gRPC Interceptors & Metadata

**Phase:** gRPC & Protocol Buffers (Days 61–65)

Adds unary server interceptors for logging and auth validation, plus a client interceptor that attaches request IDs.

## Run

```powershell
cd learn/go/day-64
go run ./cmd/grpcserver
go run ./cmd/grpcclient -addr localhost:9090
```

The server interceptor:

- Logs RPC method, request ID, status code, and duration
- Validates `authorization` metadata on `CreateBookmark` and `DeleteBookmark`
- Attaches the authenticated user to context for handlers

The client interceptor attaches `x-request-id` on every outgoing call.

## Key packages

| File | Purpose |
|------|---------|
| `internal/grpcapi/interceptors.go` | Server logging/auth + client request-id |
| `internal/grpcapi/interceptors_test.go` | Rejects unauthenticated protected RPCs |

## Tests

```powershell
go test ./...
```
