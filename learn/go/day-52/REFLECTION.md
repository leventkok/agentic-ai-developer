# Retrospective — Day 52 (JWT)

## What went well
- JWT slots in without changing handler shapes
- Env-based secret keeps keys out of git
- Algorithm allowlist blocks `none` attacks

## What was hard
- Deciding whether to drop the sessions table (left in schema, unused)
- JWT size vs opaque tokens in logs

## Next (Day 53)
- Roles (`admin` / `member`) and 403 Forbidden
