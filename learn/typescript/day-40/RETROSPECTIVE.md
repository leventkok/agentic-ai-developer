# Retrospective — Day 40 (Async, Promises & Errors capstone)

## What went well
- Typed `Promise<T>` and `async/await` chains (Days 36–37)
- Result types made expected failures explicit (Day 38)
- AbortController propagated cancellation through layers (Day 39)
- Composed retry + debounce + shutdown into one workflow (Day 40)

## What was hard
- Editor vs disk — code in IDE but not saved before `npm test`
- Knowing when to use Result vs throw vs AbortError
- Debounce with async — only the last call should win
- Exponential backoff timing in tests (needed short delays)

## Production patterns I'll reuse
- `Result<T, E>` for API clients and stores
- `retryWithBackoff` for flaky network calls
- `debounceAsync` for search/autocomplete
- `trackPromise` + `drainPending` for graceful shutdown
- Pass `AbortSignal` through every async layer

## What's next
- Node HTTP phase (Days 46–50) — same patterns on real servers
- Connect Task Tracker CLI to a real API with these patterns
