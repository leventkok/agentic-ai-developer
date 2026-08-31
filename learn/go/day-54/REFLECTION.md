# Retrospective — Day 54 (Security)

## What went well
- Rate limit middleware wraps existing handlers without invasive changes
- 429 is distinct from 401/403 — clients can retry with backoff

## What was hard
- Per-IP keys behind proxies (`X-Forwarded-For` trust boundaries)
- In-memory limiter vs distributed deployment

## Next (Day 55)
- Capstone: end-to-end auth tests + honest threat model
