# Day 45 — DOM & Browser TypeScript Capstone

**Phase:** DOM & Browser (Days 41–45)

## Run
```powershell
cd learn/typescript/day-45
npm install
npm test
npm run build
```

## What's included (Days 41–45)

| Day | Module | Topic |
|-----|--------|-------|
| 41 | `src/dom/query.ts` | Typed `querySelector`, null narrowing |
| 42 | `src/dom/events.ts` | Mouse/keyboard/submit, delegation, CustomEvent |
| 43 | `src/dom/forms.ts` | FormData, validation, form state |
| 44 | `src/api/client.ts` | Typed fetch, error unions, AbortSignal |
| 45 | `src/app/todo-app.ts` | Full typed todo app |

Tests use **jsdom** to simulate the browser in Node.

## tsconfig

`"lib": ["ES2022", "DOM"]` enables `lib.dom` types.

## Manual browser test

Open `index.html` in a browser (optional demo shell).
