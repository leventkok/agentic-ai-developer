# TASK 1 

Local: Where I write and test code on my own machine. Only me (the developer). Uses fake data and test payment keys — nothing real.

Staging: A copy of production used for testing before release. QA and the team use it. Uses fake data but a setup close to real, with test payment keys.

Production: The live app real customers use. Coffee drinkers, baristas, admins. Uses real data and live payment keys.

# TASK 2 


1. Developer works on a feature branch and commits small changes.
2. Opens a PR; CI runs build, tests, and linter.
3. After review and approval, the PR is merged into master.
4. Master is deployed to staging automatically.
5. QA smoke-tests the change on staging.
6. If it passes, tag a release (e.g. v1.0.0).
7. Deploy the tagged release to production.
8. Run a quick smoke test on prod to confirm it works.



# TASK 3 

Config item          | Local        | Staging      | Production
Payment API keys     | test keys    | test keys    | live keys
Database URL         | localhost    | staging DB   | prod DB
API base URL         | localhost    | staging URL  | prod URL
Push notification    | test keys    | test keys    | live keys
Feature flags        | all on       | test flags   | only ready ones
Log level            | debug        | debug        | error/info


# TASK 4

Deploying an untested hotfix straight to prod is risky because you skip every safety 
check — CI, QA, smoke tests — so the fix might not work or could break something else, 
and now real customers hit the bug. With payments, that means real money going wrong. 
A safer path for an urgent payment bug: reproduce it on staging first, write a minimal 
fix on a hotfix branch, and let CI pass. Deploy to staging, do a quick smoke test, then 
deploy to prod. Keep a feature flag ready so you can disable payment instantly if the 
fix still fails.