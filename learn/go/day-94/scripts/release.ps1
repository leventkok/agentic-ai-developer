# Cut a release from VERSION file
$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

$version = (Get-Content VERSION -Raw).Trim()
Write-Host "Release $version"

Write-Host "==> tests"
go test ./...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> build"
go build -o bin/api.exe ./cmd/api
go build -o bin/worker.exe ./cmd/worker

Write-Host ""
Write-Host "Tag when ready:"
Write-Host "  git tag -a $version -m `"Release $version`""
Write-Host "Notes: RELEASE_NOTES/$version.md"
Pop-Location
