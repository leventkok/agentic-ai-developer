# Verify hardening — Day 98

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

Write-Host "==> go test ./..."
go test ./...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> health probes"
go test ./internal/httpapi/... -run Health -count=1
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> observability"
& "$PSScriptRoot/verify-observability.ps1"
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> smoke load (API must be running for full test)"
if (Get-Process -Name "api" -ErrorAction SilentlyContinue) {
    & "$PSScriptRoot/load-test.ps1" -Requests 100
}

Write-Host "hardening verification passed"
Pop-Location
