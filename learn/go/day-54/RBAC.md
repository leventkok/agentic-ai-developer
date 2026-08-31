# RBAC policy — Day 53

## Roles

| Role | Permissions |
|------|-------------|
| `member` | Create bookmarks (owned). Update/delete own bookmarks only. |
| `admin` | All member permissions + modify any bookmark + bulk create |

New registrations default to `member`.

## HTTP status codes

| Code | Meaning |
|------|---------|
| **401 Unauthorized** | Missing or invalid JWT |
| **403 Forbidden** | Valid JWT but insufficient role or not the owner |

## Promoting a user to admin (dev)

```sql
UPDATE users SET role = 'admin' WHERE email = 'you@example.com';
```

Re-login to get a JWT with the updated role claim (we reload user from DB on each request).

## Centralized checks

- **Middleware:** `RequireRole(auth.RoleAdmin)` for route-level policy
- **Repository:** `canModifyBookmark(actor, bookmark)` for ownership

Keep policy in one place per layer — avoid copy-pasted role strings in handlers.
