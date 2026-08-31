# Day 63 — gRPC Server & Client

**Phase:** gRPC & Protocol Buffers (Days 61–65)

Implements the generated `BookmarkService` server and a demo client, mapping domain types and errors to gRPC.

## Run

```powershell
cd learn/go/day-63

# standalone gRPC server (port 9090)
go run ./cmd/grpcserver

# demo client (lists bookmarks; optional HTTP login + create)
go run ./cmd/grpcclient -addr localhost:9090 -http http://localhost:8080
```

Protected RPCs (`CreateBookmark`, `DeleteBookmark`) require metadata:

```
authorization: Bearer <jwt>
```

## Key packages

| Package | Purpose |
|---------|---------|
| `internal/grpcapi/mapper.go` | `domain.Bookmark` ↔ protobuf |
| `internal/grpcapi/errors.go` | Domain errors → gRPC status codes |
| `internal/grpcapi/server.go` | `BookmarkServiceServer` implementation |

## Tests

```powershell
go test ./...
```

## Regenerate stubs

```powershell
.\scripts\generate-proto.ps1
```
