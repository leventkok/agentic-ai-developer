# Verify team practices — Day 100 capstone

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

$required = @(
    "CODE_REVIEW.md",
    "CONTRIBUTING.md",
    "CHANGELOG.md",
    "ARCHITECTURE.md",
    "Makefile",
    "api/openapi.yaml",
    ".github/pull_request_template.md",
    "RELEASE_NOTES/v1.0.0.md",
    "VERSION"
)
foreach ($f in $required) {
    if (-not (Test-Path $f)) {
        Write-Error "missing required file: $f"
    }
}

$version = (Get-Content VERSION -Raw).Trim()
if ($version -ne "v1.0.0") {
    Write-Error "VERSION = $version, want v1.0.0"
}

Write-Host "==> go test ./..."
go test ./...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> go build ./..."
go build ./...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "team practices verification passed ($version)"
Pop-Location
