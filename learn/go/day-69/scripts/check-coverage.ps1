# Coverage gate — Day 69
# Measures coverage on service and sqlite packages against a minimum threshold.

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
    $profile = Join-Path $env:TEMP ("day69-" + ([guid]::NewGuid().ToString()) + ".out")
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

Write-Host "==> coverage report"
Assert-Coverage @("./internal/service") $MinServiceCoverage
Assert-Coverage @("./internal/repository/sqlite") $MinSQLiteCoverage

Write-Host "coverage gate passed"
Pop-Location
