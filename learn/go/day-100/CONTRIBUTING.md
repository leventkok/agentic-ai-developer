# Contributing

## Bootstrap

```powershell
cd learn/go/day-100
copy .env.example .env
make test
make run
```

## Makefile targets

| Target | Action |
|--------|--------|
| `make test` | Run all tests |
| `make verify` | Team practices + tests |
| `make hooks` | Install pre-commit hook |
| `make release` | Build + tag instructions |
