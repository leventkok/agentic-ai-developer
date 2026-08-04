# TASK 1 
Unit test: Tests one small function on its own. Example: formatPrice(4.5) returns "4.50 TL".
Integration test: Tests a few parts working together. Example: GET /shops/:id/menu reads from the DB and returns the menu as JSON.
E2E test: Tests the whole user flow from start to finish. Example: user opens the app, taps a shop, and sees its menu items on screen.

# TASK 2 

Unit:
- price formatter turns 4.5 into "4.50 TL"
- menu item parser reads name, price, size from a DB row
- returns empty list (not an error) when a shop has no items

Integration:
- GET /shops/:id/menu returns 200 and valid JSON for a known shop
- returns 404 when the shop id doesn't exist

E2E:
- user opens the app, picks a shop, and sees 3+ menu items
- menu screen shows a loading state, then the items

# TASK 3
Report: QA or a user finds the bug and writes it down with steps to reproduce and a screenshot. QA usually files it.
Triage: The team looks at it and decides how bad it is — is it severe, can we reproduce it, does it block the release? PM and QA mostly decide priority here.
Fix: A developer picks up the bug, finds the cause, and writes the fix on a branch.
Verify: QA re-tests the fix to confirm the bug is gone and nothing else broke.
Close: Once verified, the bug is marked closed. If it comes back, it gets reopened.

# TASK 4 

- Write unit tests for any new code
- Run the linter locally before opening a PR
- Test the happy path and at least one error path by hand
- Never commit secrets or API keys
- Update the docs when an API changes
- Check your code against the acceptance criteria
- Do a quick self-review of your own diff before asking others