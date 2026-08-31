# Optional local image scan (requires Docker Scout or Trivy installed).

param(
    [string]$Tag = "bookmarks-api:day77"
)

$ErrorActionPreference = "Stop"
if (Get-Command trivy -ErrorAction SilentlyContinue) {
    trivy image $Tag
    exit $LASTEXITCODE
}

Write-Host "trivy not installed - skip scan or install: https://aquasecurity.github.io/trivy/"
exit 0
