# Quality gate — Day 70 capstone
# Runs unit tests, coverage thresholds on critical packages, and golangci-lint.

param(
    [int]$MinServiceCoverage = 55,
    [int]$MinSQLiteCoverage = 55
)

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

function Assert-Coverage {
    param(
        [string[]]$Packages,
        [int]$MinPct
    )
    $label = ($Packages -join ", ")
    $profile = Join-Path $env:TEMP ("day70-" + ([guid]::NewGuid().ToString()) + ".out")
    go test @Packages -coverprofile=$profile
    if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }
    $output = go tool cover -func=$profile | Out-String
    if ($output -notmatch 'total:\s+\(statements\)\s+([\d.]+)%') {
        Write-Error "could not read coverage for $label"
    }
    $pct = [double]$Matches[1]
    Write-Host "$label coverage: $pct% (min $MinPct%)"
    if ($pct -lt $MinPct) {
        Write-Error "coverage below threshold for $label"
    }
}

Write-Host "==> go test ./..."
go test ./...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> coverage gates (critical paths)"
Assert-Coverage @("./internal/service") $MinServiceCoverage
Assert-Coverage @("./internal/repository/sqlite") $MinSQLiteCoverage

Write-Host "==> golangci-lint"
$lint = Get-Command golangci-lint -ErrorAction SilentlyContinue
if (-not $lint) {
    $lintPath = Join-Path $env:USERPROFILE "go\bin\golangci-lint.exe"
    if (Test-Path $lintPath) { $lint = $lintPath }
}
if (-not $lint) {
    Write-Error "golangci-lint not found; run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2"
}
& $lint run ./...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

Write-Host "==> integration tests (httpapi; container tests skip without Docker)"
go test -tags=integration ./internal/httpapi/...
if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }

if (Get-Command docker -ErrorAction SilentlyContinue) {
    Write-Host "==> optional postgres container tests"
    go test -tags=integration ./internal/test/env/...
    if ($LASTEXITCODE -ne 0) { Pop-Location; exit $LASTEXITCODE }
} else {
    Write-Host "docker not found - skipping container integration tests"
}

Write-Host "quality gate passed"
Pop-Location
