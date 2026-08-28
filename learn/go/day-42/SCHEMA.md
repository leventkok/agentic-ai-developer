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

Tracked in `schema_migrations` table.
