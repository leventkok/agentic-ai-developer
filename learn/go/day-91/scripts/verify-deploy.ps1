# Verify deploy artifacts — Day 80 capstone

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

Write-Host "==> go test"
go test ./...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> docker build"
docker build -t bookmarks-api:day80 .
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

if (Get-Command docker -ErrorAction SilentlyContinue) {
    Write-Host "==> docker compose smoke"
    docker compose up -d --build
    if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }
    Start-Sleep -Seconds 3
    try {
        $resp = Invoke-WebRequest -Uri "http://localhost:8080/bookmarks" -UseBasicParsing
        if ($resp.StatusCode -ne 200) {
            Write-Error "bookmarks status = $($resp.StatusCode)"
        }
    } finally {
        docker compose down
    }
}

Write-Host "deploy verification passed"
Pop-Location
