# TASK 1

1. [DB] Create shop, user,table migration
2. [DB] Create menu_items table migration for shops
3. [DB] Seed on sample shop + menu items for dev testing
4. [API] GET /shops/:id/menu - returns, name, price, size_options
5. [API] Return states (404 - if the shop id doesn't exist)
6. [UI] Build the menu list screen in Flutter
7. [UI] Add states (loading, error...) menu screen
8. [UI] show item name, price, and size options for each items
9. [Test] API returns 200 with valid JSON for a known shop
10. [Test] Menu screen renders 3+ items correctly
11. [DOCs] Document GET /shops/:id/menu in the READme

# TASK 2

1. git switch master && git pull          # start from an up-to-date master
2. git switch -c feature/view-menu        # create a feature branch
3. Commit DB migrations (shops, menu_items) in small commits
4. Commit the GET /shops/:id/menu API endpoint
5. Commit the Flutter menu screen (list + loading/error states)
6. git push -u origin feature/view-menu   # push the branch
7. Open a Pull Request → code review → address feedback
8. Merge into master, then delete the feature branch

# TASK 3
1. All unit and integration tests pass
2. Linter is clean (no style or format errors)
3. Code matches the story's acceptance criteria
4. PR is reviewed and approved by at least one person
5. API endpoint is documented in the README
6. No secrets or API keys are committed in the code
7. Error states (404, empty menu) are handled
8. UI works on both iOS and Android

# TASK 4

Spike: Payment gateway (iyzico vs Stripe) — we're not sure which fits KVKK and 
Turkish cards best, so we research it with a small test before Sprint 3.

Build directly: Menu API (GET /shops/:id/menu) — the design and data model are 
already clear, so we can build it right away without research.