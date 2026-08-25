# Task Tracker — TypeScript Library

Typed in-memory task store with Result-based errors.

## Install
npm install
npm run build

## Quick start
import { TaskStore, formatTask } from "./dist/index.js";

const store = new TaskStore();
const result = store.create({ title: "Learn TS" });

if (result.ok) {
  console.log(formatTask(result.value));
} else {
  console.error(result.error.code, result.error.message);
}

## API
| Export | Description |
|--------|-------------|
| TaskStore | CRUD + complete() |
| formatTask | CLI display helper |
| toPublicTask | Task → JSON-safe object |
| Result<T,E> | Typed success/failure |

## Examples
npm run example:basic
npm run example:errors

## Declaration files
After build, inspect dist/index.d.ts — that's what consumers see in IntelliSense.