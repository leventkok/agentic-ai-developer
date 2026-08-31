# gRPC Capstone — Day 65

**Phase:** gRPC & Protocol Buffers (Days 61–65)

Expose bookmark operations over HTTP and gRPC with shared services and interceptors.

## Phase recap

| Day | Focus | Key deliverable |
|-----|-------|-----------------|
| 61 | Protobuf messages | `Bookmark` message + round-trip tests |
| 62 | Service definition | `BookmarkService` RPCs + generated stubs |
| 63 | Server & client | `grpcapi` mapper/errors/server, standalone cmds |
| 64 | Interceptors | Logging, auth validation, request-id propagation |
| 65 | **Capstone** | HTTP + gRPC in one process, documented dev flow |

## Capstone checklist

### Protobuf & codegen

- [x] `api/proto/bookmarks/v1/bookmarks.proto` defines messages and unary RPCs
- [x] `scripts/generate-proto.ps1` generates Go + gRPC stubs for this module
- [x] Generated code lives in `internal/gen/bookmarksv1/`

### gRPC adapter (`internal/grpcapi`)

- [x] `mapper.go` converts `domain.Bookmark` ↔ protobuf `Bookmark`
- [x] `errors.go` maps domain errors to gRPC status codes
- [x] `server.go` implements `BookmarkServiceServer` using `service.BookmarkService`
- [x] No duplicated business logic in gRPC handlers

### Interceptors (Day 64+)

- [x] `interceptors.go` — unary server interceptor logs RPCs and validates auth metadata on protected methods
- [x] Client interceptor attaches `x-request-id` metadata
- [x] `interceptors_test.go` rejects unauthenticated create calls

### Multi-transport wiring (Day 65)

- [x] `config.Config.GRPCPort` (default `9090`)
- [x] `app.Wire` builds HTTP handler and gRPC server from shared services
- [x] `cmd/api/main.go` runs HTTP + gRPC concurrently with graceful shutdown
- [x] README documents local dev for both transports

### Auth over gRPC

- [x] Protected RPCs read JWT from metadata key `authorization` (`Bearer <token>`)
- [x] Invalid/missing tokens return `Unauthenticated`
- [x] Domain forbidden/not-found map to `PermissionDenied` / `NotFound`

### Tests

- [x] `go test ./...` passes
- [x] Mapper and status-mapping unit tests in `internal/grpcapi`

## Verify locally

```powershell
cd learn/go/day-65
go test ./...
go run ./cmd/api
# separate terminal:
go run ./cmd/grpcclient -http http://localhost:8080 -addr localhost:9090
```

## What you built

Two transports, one domain layer — the same pattern used in production systems where REST serves browsers and gRPC serves internal or high-performance clients.
