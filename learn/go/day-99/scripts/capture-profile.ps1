# Capture pprof profiles — Day 86+

param(
    [string]$BaseURL = "http://localhost:6060",
    [int]$Seconds = 10,
    [string]$OutDir = "profiles"
)

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

New-Item -ItemType Directory -Force -Path $OutDir | Out-Null

Write-Host "Capturing CPU profile (${Seconds}s)..."
Invoke-WebRequest -Uri "$BaseURL/debug/pprof/profile?seconds=$Seconds" -OutFile "$OutDir/cpu.prof"

Write-Host "Capturing heap profile..."
Invoke-WebRequest -Uri "$BaseURL/debug/pprof/heap" -OutFile "$OutDir/heap.prof"

Write-Host "Top CPU functions:"
go tool pprof -top "$OutDir/cpu.prof"

Write-Host "Profiles saved to $OutDir/"
Pop-Location
