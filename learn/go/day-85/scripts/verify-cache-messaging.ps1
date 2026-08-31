# Verify cache + messaging — Day 85 capstone

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

Write-Host "==> go test (all packages)"
go test ./...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> cache layer"
go test ./internal/repository/cached/... ./internal/cache/...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> messaging (outbox, idempotency, memory bus)"
go test ./internal/messaging/...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> service event/outbox tests"
go test ./internal/service/... -run "Enqueue|Outbox|Publish"
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

if (Test-Path "$PSScriptRoot/verify-quality.ps1") {
    Write-Host "==> quality gate"
    & "$PSScriptRoot/verify-quality.ps1"
    if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }
}

Write-Host "cache + messaging verification passed"
Pop-Location
