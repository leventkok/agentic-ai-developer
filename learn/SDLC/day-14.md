# TASK 1 

Corrective: Fixing bugs after launch. Example: menu items show the wrong price, so we patch the price formatter.
Adaptive: Adjusting to outside changes. Example: iyzico updates their payment API, so we update our integration.
Perfective: Improving performance or UX. Example: the menu screen loads slowly, so we add caching to speed it up.
Preventive: Reducing future problems. Example: refactor messy order code and add tests before it causes bugs.


# TASK 2 

Availability:
- API uptime %
- 5xx error rate on /shops/:id/menu and /orders

Performance:
- p95 response time for the menu and order endpoints
- slow DB queries (menu lookups during rush hour)

Business:
- orders per hour
- failed payment rate

Security:
- failed login attempts per user
- unusual spikes in traffic from one IP

Infrastructure:
- DB connection count
- CPU and memory usage on the API server


# TASK 3 

Detect (8:02 AM): Monitoring alerts fire — 5xx error rate on the menu API spikes. On-call engineer gets paged.
Triage (8:05 AM): Engineer confirms customers can't load menus, checks it's app-wide, marks it high severity (blocks orders during rush).
Mitigate (8:12 AM): Recent deploy is the likely cause, so they roll back to the previous stable tag to stop the bleeding.
Fix (8:30 AM): With the app stable, the developer finds the root cause on staging and prepares a proper fix.
Review (next day): Team runs a post-mortem — what broke, why monitoring caught it late, and what to add so it doesn't repeat.


# TASK 4

- A clear alert that names the service and what's failing
- A link to the runbook with rollback steps
- A list of recent deploys so I know what changed
- Error logs with a full stack trace
- A dashboard showing error rate and response times
- Who to escalate to if I'm stuck
- Access and permissions already set up so I'm not locked out mid-incident