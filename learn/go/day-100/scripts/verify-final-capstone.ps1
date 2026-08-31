# Final capstone verification — Day 100

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

$version = (Get-Content VERSION -Raw).Trim()
if ($version -ne "v1.0.0") {
    Write-Error "VERSION = $version, want v1.0.0"
}

Write-Host "==> go test ./..."
go test ./...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> health probes"
go test ./internal/httpapi/... -run Health -count=1
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> team practices"
& "$PSScriptRoot/verify-team-practices.ps1"
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> performance gate"
& "$PSScriptRoot/verify-performance.ps1"
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> cache + messaging"
& "$PSScriptRoot/verify-cache-messaging.ps1"
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "final capstone verification passed ($version)"
Pop-Location
