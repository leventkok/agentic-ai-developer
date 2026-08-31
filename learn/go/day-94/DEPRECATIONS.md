# Deprecation policy

## POST /bookmarks/bulk

- **Status:** Deprecated in v0.2.0
- **Removal:** Planned v1.0.0
- **Migration:** Use repeated `POST /bookmarks` or gRPC `BulkCreate` until removal

The HTTP handler sets `Deprecation: true` and `Sunset: 2027-01-01` response headers.
