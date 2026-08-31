# Retrospective — Day 53 (RBAC)

## What went well
- 401 vs 403 distinction is clear in middleware vs repository
- Bookmark ownership maps naturally to `user_id`
- Admin override for legacy seeded rows

## What was hard
- SQLite `ALTER TABLE` for new columns on existing data
- Composing `RequireAuth` then `RequireRole` in the right order

## Next (Day 54)
- Rate limit auth routes, dependency audit with `govulncheck`
