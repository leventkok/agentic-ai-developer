# Install git hooks from .githooks/
$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root
git config core.hooksPath .githooks
Write-Host "hooks path set to .githooks (run: make hooks)"
Pop-Location
