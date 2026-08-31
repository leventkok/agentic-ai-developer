# Run golangci-lint with shared project config.

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

$lint = Get-Command golangci-lint -ErrorAction SilentlyContinue
if (-not $lint) {
    $lintPath = Join-Path $env:USERPROFILE "go\bin\golangci-lint.exe"
    if (Test-Path $lintPath) { $lint = $lintPath }
}
if (-not $lint) {
    Write-Host "Installing golangci-lint..."
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2
    $lint = Join-Path $env:USERPROFILE "go\bin\golangci-lint.exe"
}

& $lint run ./...
$code = $LASTEXITCODE
Pop-Location
exit $code
