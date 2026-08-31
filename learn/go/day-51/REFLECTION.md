# Retrospective — Day 51 (Authentication Basics)

## What went well
- bcrypt keeps passwords out of the database as plaintext
- Session tokens in SQLite — simple opaque tokens before JWT on Day 52
- `RequireAuth` middleware composes cleanly with Go 1.22 `mux.Handle`
- Public GET routes, protected POST/PATCH/DELETE — clear auth boundary

## What was hard
- Wiring context user without import cycles (handler context helpers + middleware)
- Normalizing email (lowercase/trim) consistently in register and login
- Choosing which routes to protect vs leave public

## Production patterns I'll reuse
- **Never store plaintext passwords** — hash at the repository boundary
- **Bearer token middleware** — reusable for any protected route
- **Separate Auth repository** — auth concerns don't leak into Bookmarks repo
- **401 for missing/invalid auth** — consistent JSON problem responses

## Next (Day 52)
- Replace opaque session tokens with JWT
- Stateless token verification
