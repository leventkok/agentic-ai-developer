# Day 31 — Bookmarks API (MVP)

## Domain
Bookmarks: save URLs with a title and optional tags.

## Endpoints (planned)
| Method | Route | Description | Status |
|--------|-------|-------------|--------|
| GET | /bookmarks | List bookmarks | 200 |
| POST | /bookmarks | Create bookmark | 201 |
| GET | /bookmarks/{id} | Get one bookmark | 200 / 404 |
| DELETE | /bookmarks/{id} | Delete bookmark | 204 / 404 |

## Error format (planned)
```json
{ "code": 400, "message": "title is required" }
```

## Project layout
```
cmd/api/main.go          ← entry point (TODO)
internal/model/          ← domain types (TODO)
internal/handler/        ← HTTP handlers (TODO)
internal/store/          ← in-memory storage (TODO)
```
