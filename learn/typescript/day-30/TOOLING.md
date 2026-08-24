# Tooling decisions — Day 30 capstone

## Module system
- **ESM** (`type: "module"`) for dev and modern consumers — native `import`/`export`
- **Dual emit** (ESM + CJS) so libraries work with both `import` and `require` consumers

## Path aliases
- `@core/*` → `src/core/*` — domain types and store logic
- `@utils/*` → `src/utils/*` — presentation helpers
- Aliases live in `tsconfig.base.json` paths; `tsx --tsconfig tsconfig.esm.json` resolves them at dev time

## Build pipeline
- `build:esm` — `tsc` → `dist/esm/` then `tsc-alias` rewrites `@core/*` / `@utils/*` to relative paths for Node
- `build:cjs` — same flow → `dist/cjs/` with CommonJS emit
- `tsc` alone does not rewrite path aliases; `tsc-alias` bridges that gap for plain `node` runs

## Strict flags
- `noUncheckedIndexedAccess` — `tasks[index]` returns `Task | undefined`, forcing null checks in `toggle()` before mutating

## Dev workflow
- `npm run dev` — fast iteration via tsx, no emit step
- `npm run build` — production emit for both ESM and CJS; use before publishing or running with plain `node`
