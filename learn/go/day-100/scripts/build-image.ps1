param(
    [string]$Tag = "bookmarks-api:day77"
)

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

Write-Host "==> docker build"
docker build -t $Tag .
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "built $Tag"
Write-Host "run: docker run --rm -p 8080:8080 -p 9090:9090 -e JWT_SECRET=dev-docker-secret-at-least-32-chars $Tag"
Pop-Location
