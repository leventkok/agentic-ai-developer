# Generate Go code from .proto files (Day 65)
# Run from learn/go/day-65/

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Set-Location $root

$protoc = (Get-Command protoc -ErrorAction SilentlyContinue).Source
if (-not $protoc) {
    $wingetProtoc = Get-ChildItem -Path "$env:LOCALAPPDATA\Microsoft\WinGet\Packages" -Recurse -Filter protoc.exe -ErrorAction SilentlyContinue | Select-Object -First 1
    if ($wingetProtoc) { $protoc = $wingetProtoc.FullName }
}
if (-not $protoc -or -not (Test-Path $protoc)) {
    Write-Error "protoc not found. Install: winget install Google.Protobuf"
}

$genGo = Join-Path $env:USERPROFILE "go\bin\protoc-gen-go.exe"
if (-not (Test-Path $genGo)) {
    Write-Host "Installing protoc-gen-go..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
}

$genGoGRPC = Join-Path $env:USERPROFILE "go\bin\protoc-gen-go-grpc.exe"
if (-not (Test-Path $genGoGRPC)) {
    Write-Host "Installing protoc-gen-go-grpc..."
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
}

New-Item -ItemType Directory -Force -Path "internal\gen\bookmarksv1" | Out-Null

& $protoc `
    --proto_path=api/proto `
    --go_out=internal/gen `
    --go_opt=module=learn/go/day-65/internal/gen `
    --go-grpc_out=internal/gen `
    --go-grpc_opt=module=learn/go/day-65/internal/gen `
    api/proto/bookmarks/v1/bookmarks.proto

Write-Host "Generated internal/gen/bookmarksv1/"
