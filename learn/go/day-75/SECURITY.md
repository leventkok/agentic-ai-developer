# Security baseline — Day 55

## What we mitigate

| Threat | Mitigation |
|--------|------------|
| Plaintext passwords | bcrypt hashing |
| Stolen session DB rows | JWT stateless tokens (no session table lookup) |
| Weak/forged tokens | HS256 signature, secret from env, algorithm allowlist |
| Brute-force login | Rate limit on `/auth/register` and `/auth/login` |
| Unauthorized writes | JWT middleware on mutating routes |
| Cross-user edits | Bookmark `user_id` ownership + admin override |
| Privilege escalation on bulk ops | `RequireRole(admin)` on bulk create |
| Bad input | Validation on register/login/bookmarks |

## JWT vs server sessions (Day 52 trade-off)

| | JWT (this capstone) | DB sessions (Day 51) |
|--|---------------------|----------------------|
| Revocation | Hard — token valid until expiry | Easy — delete session row |
| DB load | Low — verify signature only | Higher — lookup each request |
| Horizontal scale | Simple — shared secret | Needs shared session store |

We still reload the user from DB after parsing JWT so role changes take effect on the next request (within token lifetime).

## Remaining risks (honest)

1. **Token revocation** — No blocklist; compromised JWT works until expiry. Mitigation: short TTL, refresh flow (future).
2. **CSRF** — Bearer tokens in headers are not CSRF-prone like cookies; cookie-based auth would need CSRF tokens.
3. **HTTPS** — Dev server is HTTP; production must terminate TLS (reverse proxy or `ListenAndServeTLS`).
4. **Secret rotation** — Changing `JWT_SECRET` invalidates all tokens; plan rotation windows.
5. **Rate limit scope** — In-memory per IP; resets on restart; use Redis for multi-instance.
6. **Admin bootstrap** — No seeded admin; promote users via DB or future admin API.

## Production checklist

- [ ] Set `ENV=production` and a random `JWT_SECRET` (32+ bytes)
- [ ] Enable HTTPS
- [ ] Run `govulncheck ./...` in CI
- [ ] Monitor 401/403/429 rates
- [ ] Document token TTL for clients
