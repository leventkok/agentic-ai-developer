# Day 51 — Authentication Basics

**Phase:** Auth & Security (Days 51–55)

## What you built

- **bcrypt** password hashing (`internal/auth/password.go`)
- **Registration** with unique email constraint
- **Login** returning opaque DB-backed session tokens
- **Auth middleware** — `Authorization: Bearer <token>` on mutating bookmark routes
- Public read routes: `GET /bookmarks`, `GET /bookmarks/{id}`

## Run

```powershell
cd learn/go/day-51
go mod tidy
go run ./cmd/api
go test ./...
```

## Try it

```powershell
# Register
curl.exe -X POST http://localhost:8080/auth/register `
  -H "Content-Type: application/json" `
  -d '{"email":"you@example.com","password":"password123"}'

# Login (save the token)
curl.exe -X POST http://localhost:8080/auth/login `
  -H "Content-Type: application/json" `
  -d '{"email":"you@example.com","password":"password123"}'

# Protected create (401 without token)
curl.exe -X POST http://localhost:8080/bookmarks `
  -H "Content-Type: application/json" `
  -d '{"title":"Go","url":"https://go.dev"}'

# With token
curl.exe -X POST http://localhost:8080/bookmarks `
  -H "Content-Type: application/json" `
  -H "Authorization: Bearer YOUR_TOKEN" `
  -d '{"title":"Go","url":"https://go.dev"}'

# Who am I?
curl.exe http://localhost:8080/auth/me `
  -H "Authorization: Bearer YOUR_TOKEN"
```

## New files

```
migrations/005_create_users.up.sql
internal/auth/password.go
internal/model/user.go
internal/repository/auth.go
internal/repository/sqlite/auth.go
internal/handler/auth.go
internal/middleware/auth.go
```

Day 52 adds JWT tokens on top of this foundation.
