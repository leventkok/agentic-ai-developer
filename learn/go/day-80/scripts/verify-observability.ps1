# Verify observability stack — Day 75 capstone

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

Write-Host "==> unit + integration tests"
go test ./...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> observability handler test"
go test ./internal/httpapi/ -run TestObservability
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> quality gate from Day 70"
powershell -ExecutionPolicy Bypass -File ".\scripts\verify-quality.ps1"
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "observability capstone passed"
Pop-Location
