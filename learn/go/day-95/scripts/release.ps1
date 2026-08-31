$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root
$version = (Get-Content VERSION -Raw).Trim()
Write-Host "Release $version — run: git tag -a $version"
go test ./...
go build -o bin/api.exe ./cmd/api
Pop-Location
