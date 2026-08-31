# Contributing

## Git workflow

1. Fork or branch from `master`
2. Commit in small, logical steps
3. Open a PR using the [template](.github/pull_request_template.md)
4. Address review feedback; squash optional per team policy

See [CODE_REVIEW.md](./CODE_REVIEW.md) for review standards.

## Bootstrap (clone → running API)

```powershell
cd learn/go/day-91
copy .env.example .env
go test ./...
go run ./cmd/api
```

HTTP: `http://localhost:8080` · gRPC: `localhost:9090` · pprof: `localhost:6060`

## Quality gates

```powershell
./scripts/verify-quality.ps1
```
