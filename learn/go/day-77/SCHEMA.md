# Bookmarks schema — Day 42

## Table: `bookmarks`

| Column | Type | Notes |
|--------|------|-------|
| `id` | INTEGER PK | Auto-increment |
| `title` | TEXT NOT NULL | Bookmark title |
| `url` | TEXT NOT NULL | Valid http/https URL |
| `tags` | TEXT NOT NULL | JSON array string |
| `created_at` | TEXT NOT NULL | ISO-ish datetime, default `now` |
| `updated_at` | TEXT NOT NULL | Updated on change (Day 43+) |

## Indexes

| Name | Column | Why |
|------|--------|-----|
| `idx_bookmarks_url` | `url` | Lookup by URL (dedup, search) |

## Migrations

| Version | Up | Down |
|---------|-----|------|
| 001 | Create `bookmarks` table | Drop table |
| 002 | Index on `url` | Drop index |
| 003 | Seed dev data | Delete seed rows |
| 004 | Create `bookmark_audit` table | Drop audit table |

Tracked in `schema_migrations` table.

## Table: `bookmark_audit`

| Column | Type | Notes |
|--------|------|-------|
| `id` | INTEGER PK | Auto-increment |
| `bookmark_id` | INTEGER FK | References `bookmarks(id)` |
| `action` | TEXT NOT NULL | e.g. `updated` |
| `detail` | TEXT | Which fields changed |
| `created_at` | TEXT NOT NULL | When audit row was written |

Updates use a transaction: change bookmark + insert audit row together.
