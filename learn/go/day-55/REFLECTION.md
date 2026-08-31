# Retrospective — Day 55 (Auth & Security capstone)

## What went well
- JWT + RBAC layered cleanly on Day 51 foundation
- Ownership checks live in the repository — handlers stay thin
- Capstone tests cover wrong password, missing token, forbidden role, cross-user edit
- `SECURITY.md` documents what we did *not* solve yet

## What was hard
- SQLite `ALTER TABLE` for roles and `user_id` on existing bookmarks
- Choosing 401 vs 403 consistently (auth vs authorization)
- Seeded bookmarks with null `user_id` needed an admin escape hatch

## Production patterns I'll reuse
- **JWT from env secret** — never hardcode signing keys
- **RequireAuth + RequireRole** — composable middleware chain
- **Rate limit sensitive routes** — login/register first
- **Reload user after JWT parse** — roles can change without re-login

## Your turn — Day 56+
Project layout & clean architecture. Use `day-55` as reference; implement the next phase yourself.
