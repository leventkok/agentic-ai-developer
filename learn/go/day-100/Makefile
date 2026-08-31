.PHONY: test lint run worker verify hooks release

test:
	go test ./...

lint:
	golangci-lint run ./...

run:
	go run ./cmd/api

worker:
	go run ./cmd/worker

verify:
	powershell -ExecutionPolicy Bypass -File scripts/verify-final-capstone.ps1

hooks:
	powershell -ExecutionPolicy Bypass -File scripts/setup-hooks.ps1

release:
	powershell -ExecutionPolicy Bypass -File scripts/release.ps1
