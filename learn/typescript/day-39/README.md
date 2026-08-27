# Day 39 — AbortController and Cancellation

**Phase:** Async, Promises & Errors (Days 36–40)

## Setup
```powershell
cd learn/typescript/day-39
npm install
```

## Tasks — which file to edit

| Task | What | File |
|------|------|------|
| **1** | Pass `AbortSignal` to async fetch | `src/services/cancellable-user.ts` |
| **2** | Propagate signal through layers | `src/services/cancellable-user.ts` |
| **3** | Detect `AbortError` vs network errors | `src/utils/abort-error.ts` |
| **4** | `withTimeout(promise, ms)` | `src/async/timeout.ts` |
| **5** | Demo + tests | `src/index.ts`, `tests/abort.test.ts` |

**Reuse from Day 38:** `src/types/errors.ts`, `src/types/result.ts`, `src/utils/safe-catch.ts`

See chat lesson for full code + logic.
