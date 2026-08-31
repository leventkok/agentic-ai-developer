# Day 55 — Auth & Security Capstone

**Phase:** Auth & Security (Days 51–55)

Final secured MVP. Builds on **Day 54** with full auth flow tests and a complete threat review.

## Phase folders

| Day | Focus | Folder |
|-----|-------|--------|
| 51 | bcrypt, sessions, protected routes | `learn/go/day-51` |
| 52 | JWT replaces sessions | `learn/go/day-52` |
| 53 | RBAC + bookmark ownership | `learn/go/day-53` |
| 54 | Rate limit + security hygiene | `learn/go/day-54` |
| 55 | **Capstone** — tests + threat model | `learn/go/day-55` |

## Run

```powershell
cd learn/go/day-55
go mod tidy
go run ./cmd/api
go test ./...
```

Copy `.env.example` to `.env` and set a strong `JWT_SECRET` before production.

## Auth for API clients

### 1. Register

```powershell
curl.exe -X POST http://localhost:8080/auth/register `
  -H "Content-Type: application/json" `
  -d '{"email":"you@example.com","password":"password123"}'
```

### 2. Login (returns JWT)

```powershell
curl.exe -X POST http://localhost:8080/auth/login `
  -H "Content-Type: application/json" `
  -d '{"email":"you@example.com","password":"password123"}'
```

Response:

```json
{"token":"eyJ...","user":{"id":1,"email":"you@example.com","role":"member",...}}
```

### 3. Send token on protected routes

```powershell
curl.exe -X POST http://localhost:8080/bookmarks `
  -H "Content-Type: application/json" `
  -H "Authorization: Bearer YOUR_JWT" `
  -d '{"title":"Go","url":"https://go.dev"}'
```

## Route policy

| Route | Auth | Role / rule |
|-------|------|-------------|
| `GET /bookmarks` | Public | — |
| `GET /bookmarks/{id}` | Public | — |
| `POST /auth/register`, `POST /auth/login` | Public | Rate limited |
| `GET /auth/me` | JWT | — |
| `POST /bookmarks` | JWT | Sets `user_id` to caller |
| `PATCH/DELETE /bookmarks/{id}` | JWT | Owner or `admin` |
| `POST /bookmarks/bulk` | JWT | `admin` only |

Seeded bookmarks have no owner — only `admin` can modify them.

## Dependency audit

```powershell
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
```

See `SECURITY.md` for remaining risks (token revocation, CSRF scope, HTTPS).

---

**Your turn next:** Day 56+ (project layout & architecture) — you'll write the code yourself. Use this folder as the reference implementation.
