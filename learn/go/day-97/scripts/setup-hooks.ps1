$root = Split-Path -Parent $PSScriptRoot
Push-Location $root
git config core.hooksPath .githooks
Write-Host "hooks installed"
Pop-Location
