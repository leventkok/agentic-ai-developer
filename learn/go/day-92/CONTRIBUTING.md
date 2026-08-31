# Contributing

## Git workflow

1. Fork or branch from `master`
2. Commit in small, logical steps
3. Open a PR using the [template](.github/pull_request_template.md)
4. Address review feedback

See [CODE_REVIEW.md](./CODE_REVIEW.md).

## Bootstrap

```powershell
cd learn/go/day-92
copy .env.example .env
go test ./...
go run ./cmd/api
```

## Quality gates

```powershell
./scripts/verify-quality.ps1
```
