
# TASK 1 

1. Correctness: Does the code do what the acceptance criteria describe, including edge cases?
2. Tests: Are there tests, and do they also cover failure paths (not just the happy path)?
3. Security: No secrets in the code, inputs are validated, and queries are safe from SQL injection.
4. Naming: Are variables and functions named clearly and consistently?
5. Error handling: Are errors caught, logged, and returned with the right status code?
6. Performance: No N+1 queries or unnecessary loops that slow things down.
7. Readability: Could a junior developer understand this code 6 months from now?
8. Docs: Are API changes documented in the README?
9. Comparison logic: Are comparisons correct (=== vs =), and are booleans handled properly?
10. Scope: Does the PR stay focused on one feature, without unrelated changes?

# TASK 2 

Kind comment: "Nice job handling the 404 case when no menu items are found — good to think about empty results early."

Critical comment (blocking): "The query builds SQL with a string template (shop_id = ${shopId}), which is open to SQL injection. Please use a parameterized query before merge. Also, the console.log prints STRIPE_SECRET — that secret must be removed from the logs."


# TASK 3

Thanks for the suggestion. Can you share what problems you're seeing with REST here — is it about flexibility or something specific? For this project we chose REST because it fits our mini-spec and the team already knows it well, so it keeps the MVP fast to build. If we hit real limitations later, I'm happy to revisit GraphQL for those parts. For now I'd prefer to keep REST so we don't slow down the sprint.




# TASK 4 

1. What changed and why — a short summary of the feature
2. Link to the user story or ticket (e.g. "view menu" story)
3. How to test it — steps for the reviewer to run it locally
4. Screenshots or a short video for UI changes (the Flutter menu screen)
5. Which acceptance criteria this PR covers
6. Known limitations or things left for a later PR
7. Any new environment variables or migrations the reviewer must run
8. Type of change — new feature, bug fix, or refactor
9. Checklist confirming tests pass and the linter is clean