# Protocol Buffers — Day 61 Guide

**New phase:** gRPC & Protocol Buffers (Days 61–65)

## What is protobuf? (plain language)

JSON is text humans read easily. **Protobuf** is a **binary format** defined by a **schema file** (`.proto`). You write the schema once; tools generate Go structs and encode/decode code.

```
.proto file  →  protoc  →  bookmarks.pb.go  →  your Go program
(schema)       (tool)      (generated)         (marshal/unmarshal)
```

## Files in this project

| Path | Purpose |
|------|---------|
| `api/proto/bookmarks/v1/bookmarks.proto` | Schema — **source of truth** |
| `internal/gen/bookmarksv1/bookmarks.pb.go` | Generated — **do not edit by hand** |
| `scripts/generate-proto.ps1` | Regenerate after `.proto` changes |
| `cmd/protodemo/main.go` | Your marshal/unmarshal exercise |

## Install tools (one time)

```powershell
winget install Google.Protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Restart terminal so `protoc` is on PATH.

## Regenerate Go code

```powershell
cd learn/go/day-61
.\scripts\generate-proto.ps1
```

## Your tasks today

1. Read `bookmarks.proto` — notice `message`, field numbers, `repeated`, `optional`
2. Open generated `bookmarks.pb.go` — see struct tags and `GetTitle()` methods
3. Implement `cmd/protodemo/main.go` (remove `panic`, use `proto.Marshal` / `proto.Unmarshal`)
4. Un-skip tests in `bookmarks_test.go`
5. Run:

```powershell
go test ./internal/gen/bookmarksv1/...
go run ./cmd/protodemo
```

## Field numbers matter

In `.proto`, `string title = 1` — the `1` is permanent on the wire. Never reuse numbers for different fields.

## Proto vs domain

| Layer | Package | Today |
|-------|---------|-------|
| Domain | `internal/domain` | Business rules (Day 57–59) |
| Proto | `api/proto/...` | Wire format for RPC (Day 61+) |

Mapping between them comes on later days — not today.

## JSON vs protobuf (size)

Protobuf is usually smaller and faster to parse — good for service-to-service calls. REST/JSON stays for browsers and humans.
