# Day 35 — Folder layout

```
day-35/
├── package.json          ← main, types, bin.tasks
├── tsconfig.json
├── README.md
├── RETROSPECTIVE.md
├── FOLDER.md
├── tests/
│   └── store.test.ts     ← smoke tests
└── src/
    ├── index.ts          ← library entry (re-exports api)
    ├── api/
    │   └── index.ts      ← public exports only
    ├── types/
    │   ├── task.ts
    │   └── result.ts
    ├── core/
    │   ├── store.ts
    │   └── errors.ts
    ├── utils/
    │   ├── format.ts
    │   ├── parse.ts
    │   └── index.ts
    └── cli/
        └── app.ts        ← CLI entry (bin.tasks)
```

## Rules

- Everything lives under `src/` — never put code at project root
- `src/index.ts` → library entry (`main` / `types`)
- `src/cli/app.ts` → CLI entry (`bin.tasks`), not exported from `api/`
- `tests/` sits beside `src/`, not inside it
- Build output goes to `dist/` (gitignored)
