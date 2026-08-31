# Package Diagram — Day 59

Sketch import arrows to verify dependency direction. **Update this diagram** if your packages differ.

## Allowed dependencies (inward only)

```mermaid
flowchart TB
    subgraph outer["Transport & infra"]
        cmd["cmd/api"]
        app["internal/app"]
        httpapi["internal/httpapi"]
        middleware["internal/middleware"]
        auth["internal/auth"]
        config["internal/config"]
        db["internal/db"]
    end

    subgraph application["Application"]
        service["internal/service"]
    end

    subgraph core["Core"]
        domain["internal/domain"]
    end

    subgraph persistence["Persistence"]
        repo_iface["internal/repository"]
        sqlite["internal/repository/sqlite"]
        memory["internal/repository/memory"]
    end

    cmd --> app
    app --> httpapi
    app --> sqlite
    app --> auth
    app --> config
    app --> db

    httpapi --> service
    httpapi --> domain
    httpapi --> middleware
    httpapi --> ctxkey["internal/ctxkey"]

    middleware --> domain
    middleware --> repository

    service --> domain
    service --> repository

    sqlite --> domain
    sqlite --> db
    memory --> domain

    repo_iface --> domain
```

## Forbidden edges (should NOT exist)

| From | To | Why |
|------|-----|-----|
| `domain` | `httpapi`, `service`, `repository`, `net/http` | Dependency rule |
| `service` | `httpapi`, `net/http` | Services are transport-agnostic |
| `httpapi` | `repository/sqlite` | Handlers talk to services, not DB |
| `repository` | `httpapi`, `service` | Persistence doesn't know use cases |

## How to verify

```powershell
# Quick manual checks
rg "net/http" internal/domain/
rg "database/sql" internal/domain/
rg "repository/sqlite" internal/httpapi/
rg "repository/sqlite" internal/service/

# Automated (after implementing layers_test.go)
go test ./internal/architecture/...
```

## Import cycle detection

```powershell
go list -f '{{.ImportPath}} -> {{join .Imports ", "}}' ./internal/...
```

If Go reports an import cycle during `go build`, draw the packages involved and break the cycle (usually extract shared types to `domain` or `ctxkey`).
