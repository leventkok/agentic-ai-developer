# Task Tracker — TypeScript MVP

Typed in-memory task store with JSON persistence, Result-based errors, and a CLI.

## Install

```bash
npm install
npm run build
```

## CLI

```bash
npm start -- list
npm start -- add "Learn TypeScript"
npm start -- done <id>
npm start -- delete <id>
```

After build, the `tasks` bin is available:

```bash
npx tasks list
```

Tasks persist in `tasks.json` in the working directory.

## Library

```typescript
import { TaskStore, formatTask } from "./dist/index.js";

const store = new TaskStore();
const result = store.create({ title: "Learn TS" });

if (result.ok) {
  console.log(formatTask(result.value));
} else {
  console.error(result.error.code, result.error.message);
}
```

## API

| Export | Description |
|--------|-------------|
| `TaskStore` | CRUD + `complete()` with JSON file persistence |
| `formatTask` | CLI display helper |
| `toPublicTask` | Task → JSON-safe object |
| `Result<T, E>` | Typed success/failure |
| `StoreError` | `{ code, message }` error shape |

## Tests

```bash
npm test
```

## Declaration files

After `npm run build`, inspect `dist/index.d.ts` — that is what consumers see in IntelliSense.
