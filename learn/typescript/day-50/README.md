# Day 50 — Node.js TypeScript Capstone

**Phase:** Node.js TypeScript (Days 46–50)

## Run
```powershell
cd learn/typescript/day-50
npm install
npm test
npm run dev
```

Set env vars (see `.env.example`):
```powershell
$env:PORT="3000"; $env:DATA_FILE="./data/notes.json"; npm run dev
```

## Phase summary

| Day | Package | Topic |
|-----|---------|-------|
| 46 | project setup | `package.json`, `nodenext`, npm scripts |
| 47 | `src/storage/` | `fs/promises`, JSON config file |
| 48 | `src/server/` | Typed HTTP handlers, JSON bodies |
| 49 | `src/config/` | Typed `process.env`, fail-fast |
| 50 | all layers | Structured Notes API + tests |

## Structure

```
src/
  config/     ← typed env (Day 49)
  storage/    ← fs persistence (Day 47)
  services/   ← business logic
  handlers/   ← HTTP parsing (Day 48)
  server/     ← routing
  types/      ← Result + AppError
```

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | /health | Health check |
| GET | /notes | List notes |
| POST | /notes | Create note |
| GET | /notes/:id | Get note |
| PATCH | /notes/:id | Update note |
| DELETE | /notes/:id | Delete note |
