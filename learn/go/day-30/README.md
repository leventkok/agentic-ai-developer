# Day 30 — Notes API

## Run
go run .

## Endpoints
| Method | Route | Description |
| GET | /notes | List notes |
| POST | /notes | Create note |
| DELETE | /notes/{id} | Delete note |

## Examples
curl http://localhost:8080/notes
curl -X POST http://localhost:8080/notes -H "Content-Type: application/json" -d "{\"title\":\"Hi\",\"content\":\"test\"}"
curl -X DELETE http://localhost:8080/notes/1