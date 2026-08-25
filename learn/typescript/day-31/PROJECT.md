# Task Tracker — TypeScript MVP

## Goal
A typed task-tracker library with a CLI. Users can add, list, complete, and delete tasks.

## Public API (Week 31–35)
| Export | Type | Purpose |
|--------|------|---------|
| `Task` | type | Domain entity |
| `CreateTaskInput` | type | Input for new tasks |
| `UpdateTaskInput` | type | Partial update input |
| `PublicTask` | type | Serializable task for JSON output |
| `TaskStore` | class | (Day 32) In-memory CRUD |
| `formatTask` | function | (Day 33) Display helper |

## Success criteria
- [ ] Strict TypeScript, no `any`
- [ ] Exported types match runtime behavior
- [ ] CLI runs: add, list, done, delete
- [ ] Tests cover happy + unhappy paths

## Non-goals (out of scope)
- Database persistence
- User authentication
- Web UI / REST API
- Multi-user / sync

## Architecture
src/ types/ ← domain types (Day 31) core/ ← store logic (Day 32) utils/ ← formatters, parsers (Day 33) cli/ ← entry point (Day 34)

