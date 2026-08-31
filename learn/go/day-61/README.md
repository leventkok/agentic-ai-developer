# Day 61 — Protocol Buffers Basics

**Phase:** gRPC & Protocol Buffers (Days 61–65)

Builds on Day 60 architecture. **You write the code** — demo and tests are scaffolded.

> Go bilmiyorsan sorun değil — bu gün HTTP API yazmıyorsun; sadece `.proto` şeması ve serialize pratiği.

## Bugün ne öğreneceksin?

1. `.proto` dosyası yazmak (mesaj tanımları)
2. `protoc` ile Go kodu üretmek
3. `proto.Marshal` / `proto.Unmarshal` ile byte'a çevirmek

## Adım adım (senin işin)

| Adım | Dosya | Ne yap |
|------|-------|--------|
| 1 | `api/proto/bookmarks/v1/bookmarks.proto` | Oku; alan numaralarını anla |
| 2 | `internal/gen/bookmarksv1/bookmarks.pb.go` | Generate edilmiş kodu incele (elle düzenleme) |
| 3 | `cmd/protodemo/main.go` | Marshal/unmarshal demo — `panic`'i kaldır |
| 4 | `internal/gen/bookmarksv1/bookmarks_test.go` | Skip'li testleri aç |
| 5 | `PROTOBUF.md` | Checklist |

## Kurulum (ilk kez)

```powershell
winget install Google.Protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Terminali yeniden aç.

## Çalıştır

```powershell
cd learn/go/day-61
.\scripts\generate-proto.ps1      # .proto değiştirince
go test ./internal/gen/bookmarksv1/...
go run ./cmd/protodemo
go test ./...                      # mevcut HTTP API hâlâ çalışır
go run ./cmd/api                   # Day 60 API aynı
```

## Önemli

- `internal/gen/` → **generated** — el ile yazma
- `api/proto/` → **senin schema'n** — burayı düzenle
- gRPC server **bugün yok** — Gün 62+

---

Detaylı rehber: `PROTOBUF.md`. Takılırsan **"can u fix"** de.
