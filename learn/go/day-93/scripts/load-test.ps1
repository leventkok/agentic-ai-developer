# Simple load test for GET /bookmarks — Day 90

param(
    [string]$BaseURL = "http://localhost:8080",
    [int]$Requests = 200
)

$ErrorActionPreference = "Stop"
$sw = [System.Diagnostics.Stopwatch]::StartNew()
$ok = 0
for ($i = 0; $i -lt $Requests; $i++) {
    $resp = Invoke-WebRequest -Uri "$BaseURL/bookmarks" -UseBasicParsing
    if ($resp.StatusCode -eq 200) { $ok++ }
}
$sw.Stop()
$ms = $sw.ElapsedMilliseconds
Write-Host "completed $ok/$Requests in ${ms}ms ($([math]::Round($Requests * 1000.0 / [math]::Max($ms,1), 1)) req/s)"
