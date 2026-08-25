# Retrospective — Day 35

## What went well
- Designed a clear public API up front (`TaskStore`, `Result<T,E>`, domain types) and kept it stable from Day 31 through Day 35
- Split the project into layers: `types/` → `core/` → `utils/` → `cli/`, with a thin `api/index.ts` for exports
- Used `Result<T, E>` instead of throwing — callers must handle errors explicitly, which made the CLI and tests straightforward
- Added JSON file persistence so the CLI works across separate `npm start` runs (not just in-memory)
- Smoke tests cover both happy path (create + list) and error path (`EMPTY_TITLE`)
- Package entry points are wired correctly: `main`, `types`, and `bin.tasks` after `npm run build`

## What was hard
- Keeping code under `src/` instead of duplicating folders at the project root (Day 34 had `types/`, `core/`, and `utils/` in two places)
- Editor vs disk — files looked saved in the IDE but `tsc` couldn't find them until everything was actually written to disk
- Copying `{ ... }` placeholders literally from instructions instead of real implementation (Day 33 store)
- Exporting Day 32+ modules too early in Day 31's `index.ts` before those files existed
- Each CLI run is a new Node process — without persistence, `list` looked empty after `add` until `tasks.json` was added
- PowerShell quirks: `curl` is an alias for `Invoke-WebRequest`; use `curl.exe` or `Invoke-RestMethod` when testing APIs

## What I'd change next time
- Lock the folder layout on Day 31 (`FOLDER.md` + empty scaffold) so duplicates never appear at the root
- Add persistence earlier if the project has a CLI — don't assume in-memory state survives between runs
- Write one integration test that runs the CLI end-to-end (`add` then `list` in the same script) sooner
- Add a `prebuild` lint or a simple script that fails if `.ts` files exist outside `src/` or `tests/`
- Consider a real database or SQLite only when the brief actually needs it — file JSON was enough for this MVP

## Where types caught bugs
- `Result<T, E>` forced every `store.create()` call site to check `result.ok` before using `result.value` — no silent undefined access
- `StoreError` with a `code` union (`EMPTY_TITLE`, `NOT_FOUND`, etc.) made test assertions precise instead of comparing loose strings
- `CreateTaskInput` vs `UpdateTaskInput` (with `Partial`) prevented passing wrong shapes into update logic
- `PublicTask` vs `Task` separated JSON-safe output from internal `Date` fields — avoids accidental serialization bugs
- Strict `tsconfig` flagged missing imports and wrong `.js` extension paths before runtime
