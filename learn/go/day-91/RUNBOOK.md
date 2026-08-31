# Runbook — Bookmarks API (Day 75)

Short on-call steps when things go wrong.

## Error rate spike

1. Check structured logs for `level=ERROR` and note `request_id`.
2. Open `/metrics` — compare 5xx counts on `http_requests_total`.
3. If DB errors cluster, check SQLite file path and disk space.
4. If breaker open (slow lists), pause traffic briefly; verify DB connectivity.
5. Roll back latest deploy if errors started after a release.

## High latency

1. Inspect trace spans (stdout exporter in dev) for slow repository calls.
2. Check `http_request_duration_seconds` histogram buckets.
3. Confirm `LIST_TIMEOUT_MS` is appropriate for dataset size.

## Auth failures

1. Never log JWT secrets or passwords.
2. Verify `JWT_SECRET` matches across instances.
3. Check clock skew if tokens expire immediately.

## Health checks

```powershell
curl http://localhost:8080/bookmarks
curl http://localhost:8080/metrics
```

## Escalation

- Capture request_id + timestamp + user impact.
- Attach last 50 JSON log lines (redact tokens).
- File incident with repro steps.
