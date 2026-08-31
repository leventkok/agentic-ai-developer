# Day 52 — JWT and Sessions

**Phase:** Auth & Security (Days 51–55)

Replace Day 51 opaque session tokens with **signed JWTs**.

## What you built

- `internal/auth/jwt.go` — issue and parse HS256 tokens
- Login returns a JWT; middleware validates signature + expiry
- `JWT_SECRET` and `JWT_TTL_HOURS` from environment
- `JWT.md` — trade-offs vs server-side sessions

## Run

```powershell
cd learn/go/day-52
go mod tidy
go run ./cmd/api
go test ./...
```

## Try it

```powershell
curl.exe -X POST http://localhost:8080/auth/login `
  -H "Content-Type: application/json" `
  -d '{"email":"you@example.com","password":"password123"}'
```

Use the `token` value as `Authorization: Bearer <token>` on protected routes.

**Next:** Day 53 adds roles and authorization (403).
