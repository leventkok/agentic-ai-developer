# JWT vs server-side sessions

Day 51 used **opaque DB-backed session tokens**. Day 52 switches to **JWTs**.

## Server-side sessions (Day 51)

| Pros | Cons |
|------|------|
| Easy revocation — delete the row | DB lookup on every request |
| Small token on the wire | Harder to scale horizontally without shared store |
| Server controls lifetime tightly | More moving parts (sessions table) |

## JWTs (Day 52)

| Pros | Cons |
|------|------|
| Stateless — verify signature only | Revocation is hard until token expires |
| Scales horizontally with shared secret | Token size larger than opaque ID |
| Claims travel with the token | Must guard signing key carefully |

## What we do here

- Sign with **HS256** and `JWT_SECRET` from env (never in source control)
- Reject unexpected algorithms (no `none`)
- Reload user from DB after parse so password changes can still matter for future flows
- Sessions table remains from migration 005 but is **unused** — optional cleanup later

Day 53 adds roles; Day 54 adds rate limiting; Day 55 is the full capstone.
