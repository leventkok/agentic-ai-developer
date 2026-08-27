# Day 40 — Async Integration Patterns (Capstone)

**Phase:** Async, Promises & Errors (Days 36–40)

Combines Days 36–39:
- Promise types + `async/await` (36–37)
- Result + typed errors (38)
- AbortController (39)
- **Today:** retry, debounce, graceful shutdown, composed workflow

## Run
```powershell
cd learn/typescript/day-40
npm install
npm start
npm test
```

## Patterns

| Pattern | File | Purpose |
|---------|------|---------|
| Retry + backoff | `src/async/retry.ts` | Transient failures (503, timeout) |
| Debounce | `src/async/debounce.ts` | Search inputs — one in-flight request |
| Graceful shutdown | `src/async/shutdown.ts` | Drain pending on SIGINT |
| Composed workflow | `src/services/fetch-workflow.ts` | Result + retry + abort + track |

## Manual SIGINT test
```powershell
npm start
# Ctrl+C → "SIGINT received — draining pending work..."
```
