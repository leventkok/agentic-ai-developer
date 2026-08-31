# Day 53 — Authorization and RBAC

**Phase:** Auth & Security (Days 51–55)

Builds on Day 52 JWT auth. Adds **roles** and **resource ownership**.

## What you built

- Migration `006` — `users.role` (`admin` / `member`), `bookmarks.user_id`
- `RequireRole` middleware — returns **403** when role is insufficient
- Ownership checks — members edit only their bookmarks; admins edit any
- `POST /bookmarks/bulk` — **admin only**

## Run

```powershell
cd learn/go/day-53
go mod tidy
go run ./cmd/api
go test ./...
```

## Route policy

| Route | Auth | Rule |
|-------|------|------|
| `POST /bookmarks` | JWT | Sets `user_id` to caller |
| `PATCH/DELETE /bookmarks/{id}` | JWT | Owner or `admin` |
| `POST /bookmarks/bulk` | JWT | `admin` only |

Seeded bookmarks have no owner — only admins can modify them.

See `RBAC.md` for role definitions.

**Next:** Day 54 adds rate limiting and security hygiene.
