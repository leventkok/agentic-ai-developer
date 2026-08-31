# Day 62 — Defining gRPC Services (Complete)

**Phase:** gRPC & Protocol Buffers (Days 61–65)

Adds unary RPC definitions to the protobuf schema and generates server/client interfaces.

## What changed

- `bookmarks.proto` — `BookmarkService` with 4 unary RPCs
- Generated `bookmarks_grpc.pb.go` — server/client interfaces
- `scripts/generate-proto.ps1` — includes `protoc-gen-go-grpc`

## Regenerate

```powershell
cd learn/go/day-62
.\scripts\generate-proto.ps1
```

## Next

Day 63 — implement server and client (`learn/go/day-63/`).
