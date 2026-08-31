.PHONY: test lint run verify release

test:
	go test ./...

lint:
	golangci-lint run ./...

run:
	go run ./cmd/api

verify:
	powershell -ExecutionPolicy Bypass -File scripts/verify-quality.ps1

release:
	powershell -ExecutionPolicy Bypass -File scripts/release.ps1
