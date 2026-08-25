# Day 34 — Folder layout

```
day-34/
├── package.json
├── tsconfig.json
├── README.md
├── tests/
│   └── store.test.ts       ← smoke tests (YOU write)
└── src/
    ├── index.ts            ← re-exports public API
    ├── api/
    │   └── index.ts        ← public exports only
    ├── types/
    │   ├── task.ts         ← domain types (refactor with Partial)
    │   └── result.ts       ← Result<T,E>
    ├── core/
    │   ├── store.ts        ← TaskStore (from Day 33)
    │   └── errors.ts       ← StoreError
    ├── utils/
    │   ├── format.ts       ← formatTask, toPublicTask
    │   ├── parse.ts        ← parseCommand (YOU write)
    │   └── index.ts        ← re-export utils
    └── cli/
        └── app.ts          ← CLI entry (YOU write)
```

## Rules
- Everything lives under `src/` — never put code at project root
- `src/index.ts` → library entry
- `src/cli/app.ts` → CLI entry (not exported from api/)
- `tests/` sits beside `src/`, not inside it
