# Performance verification — Day 90 capstone

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

Write-Host "==> go test ./..."
go test ./...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> benchmark regression guard"
go test ./internal/repository/sqlite/... -run TestListBenchmarkRegressionGuard -count=1
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> benchmem sample"
go test -bench=BenchmarkStoreList -benchmem ./internal/repository/sqlite/... -count=1

Write-Host "==> explain query plan"
go test ./internal/db/... -run TestExplain -count=1
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

if (Test-Path "$PSScriptRoot/verify-quality.ps1") {
    Write-Host "==> quality gate"
    & "$PSScriptRoot/verify-quality.ps1"
    if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }
}

Write-Host "performance verification passed"
Pop-Location
