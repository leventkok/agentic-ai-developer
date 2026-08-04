# TASK 1 

Database
[ ] Migrations run on production
[ ] Dev seed data removed from prod

Config
[ ] Production environment variables set (no dev keys in prod)
[ ] iyzico/Stripe keys point to live accounts, not test

Security
[ ] HTTPS enabled on all endpoints
[ ] Secrets stored in a vault, not in the code
[ ] KVKK-related user data handled correctly

Feature flags
[ ] Payment enabled only after it's fully tested

Testing
[ ] All unit, integration, and E2E tests pass
[ ] QA sign-off received

Smoke test
[ ] View menu works after deploy
[ ] Cart and payment work end to end
[ ] Shop receives the order notification

Rollback
[ ] Previous version tagged in Git
[ ] Rollback steps written down

Monitoring
[ ] Error tracking is live
[ ] Alerts set for failed payments

Comms
[ ] Release notes published
[ ] Coffee shops and support team notified


# TASK 2 

CI (Continuous Integration) runs automated checks every time someone pushes code or 
opens a PR, so nobody has to remember to run them by hand. On each push it builds the 
project, runs the tests, and runs the linter. This supports my Day 9 coding DoD (tests 
pass, linter clean) by enforcing it automatically instead of trusting people to check. 
It also backs up the Day 10 review gate — a reviewer can trust the code already passed 
the basics before they even look. If CI fails before merge, the PR is blocked until the 
problem is fixed, so broken code never reaches master.


# TASK 3 

Scenario: after deploy, customers are charged but orders don't reach shops.

Detection: monitoring alerts show a spike in failed orders, and shops/customers 
report missing orders.

Immediate action: turn off the payment feature flag so no new customers get charged 
while the bug is live.

Git rollback: revert to the previous stable tag (e.g. v0.9.0) and redeploy it, or 
revert the bad commit and push a fixed build.

Communication: tell the coffee shops the app is paused, and message affected customers 
that refunds are on the way.

Post-incident: refund the charged customers, find the root cause, write a fix, and add 
a test so it can't happen again.

# TASK 4 

# RotaCoffee v1.0.0 — MVP Pre-Order

## Added
- View a shop's menu with name, price, and size options
- Add coffees to a cart before ordering
- Pay through the app (iyzico)
- Shop receives an order notification after payment
- Barista can mark an order as "Ready"

## Known limitations
- No favorites, scheduled pickup, or promotions yet
- Requires an internet connection

## Technical
- Flutter mobile app (iOS + Android)
- REST API with PostgreSQL