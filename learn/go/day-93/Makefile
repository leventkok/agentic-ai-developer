.PHONY: test lint run migrate verify hooks

test:
	go test ./...

lint:
	golangci-lint run ./...

run:
	go run ./cmd/api

worker:
	go run ./cmd/worker

migrate:
	@echo "Migrations apply automatically on API startup (internal/db.RunMigrations)"

verify:
	go test ./...
	powershell -ExecutionPolicy Bypass -File scripts/verify-quality.ps1

hooks:
	powershell -ExecutionPolicy Bypass -File scripts/setup-hooks.ps1
