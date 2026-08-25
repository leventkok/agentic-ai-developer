# Day 35 — Bookmarks API MVP

## Run
```powershell
cd learn/go/day-35
go run ./cmd/api
```

## Endpoints
| Method | Route | Status | Description |
|--------|-------|--------|-------------|
| GET | /bookmarks | 200 | List all bookmarks |
| POST | /bookmarks | 201 | Create bookmark |
| GET | /bookmarks/{id} | 200 / 404 | Get one bookmark |
| PATCH | /bookmarks/{id} | 200 / 404 | Partial update |
| DELETE | /bookmarks/{id} | 204 / 404 | Delete bookmark |

## Error format
```json
{ "code": 400, "message": "title is required" }
```

## Manual test checklist
- [ ] POST valid bookmark → 201
- [ ] POST empty title → 400
- [ ] POST invalid URL → 400
- [ ] GET unknown id → 404
- [ ] GET id=abc → 400
- [ ] PATCH title → 200
- [ ] DELETE existing → 204
- [ ] DELETE again → 404

## Tests
```powershell
go test ./...
go vet ./...
```

## Architecture
```
cmd/api/          → entry point + routes + middleware
internal/handler/ → HTTP layer (thin)
internal/store/   → repository interface + memory impl
internal/validation/ → input rules
internal/model/   → domain + DTO types
```
