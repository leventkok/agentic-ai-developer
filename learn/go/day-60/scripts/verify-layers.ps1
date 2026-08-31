# Verify layer boundaries — Day 60 capstone
# Run from learn/go/day-60/

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Set-Location $root

Write-Host "=== Domain must not import outer layers ===" -ForegroundColor Cyan
$forbidden = 'net/http|database/sql|learn/go/day-60/internal/service|learn/go/day-60/internal/httpapi|learn/go/day-60/internal/repository'
$domainHits = rg "^\s+`"($forbidden)`"" internal/domain/ --glob "*.go" 2>$null
if ($domainHits) {
    Write-Host "FAIL: forbidden imports in domain:" -ForegroundColor Red
    $domainHits
    exit 1
}
Write-Host "OK" -ForegroundColor Green

Write-Host "=== Service must not import httpapi or net/http ===" -ForegroundColor Cyan
$svcHits = rg '^\s+"(net/http|learn/go/day-60/internal/httpapi)"' internal/service/ --glob "*.go" --glob "!*_test.go" 2>$null
if ($svcHits) {
    Write-Host "FAIL:" -ForegroundColor Red
    $svcHits
    exit 1
}
Write-Host "OK" -ForegroundColor Green

Write-Host "=== httpapi must not import sqlite directly ===" -ForegroundColor Cyan
$apiHits = rg '^\s+"learn/go/day-60/internal/repository/sqlite"' internal/httpapi/ --glob "*.go" 2>$null
if ($apiHits) {
    Write-Host "FAIL:" -ForegroundColor Red
    $apiHits
    exit 1
}
Write-Host "OK" -ForegroundColor Green

Write-Host "=== Full test suite ===" -ForegroundColor Cyan
go test ./...
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "=== Build ===" -ForegroundColor Cyan
go build ./cmd/api
if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }

Write-Host "`nAll architecture checks passed." -ForegroundColor Green
