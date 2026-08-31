# Architecture — Bookmarks API

> Day 93 adds Makefile DX. Same layered design as Day 92.

## Overview

```
Client → HTTP/gRPC → middleware → service → domain → repository
```

## Makefile targets

`make test`, `make lint`, `make run`, `make verify`, `make hooks`

See [CONTRIBUTING.md](./CONTRIBUTING.md) for bootstrap.
