# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder
WORKDIR /src

# Cache module downloads in a separate layer.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-s -w" \
    -trimpath \
    -o /out/api ./cmd/api

FROM alpine:3.20
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /out/api /app/api
COPY migrations ./migrations

ENV PORT=8080
ENV GRPC_PORT=9090
ENV ENV=production
ENV DB_PATH=/tmp/bookmarks.db

EXPOSE 8080 9090
USER 65532:65532
ENTRYPOINT ["/app/api"]
