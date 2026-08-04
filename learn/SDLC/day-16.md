# TASK 1 

Feature: Pre-order and pay for coffee
User: Coffee drinker
Value: Skip the morning queue by ordering and paying before arrival


# TASK 2 

# Pre-order and Pay — Full SDLC Packet

## Requirements
- Story: As a coffee drinker, I want to pre-order and pay before I arrive, so that I can skip the morning queue.
- AC 1: Payment must succeed before the order is sent to the shop.
- AC 2: After paying, I see an order confirmation with an estimated pickup time within 3 seconds.
- AC 3: If payment fails, no order is created and I see an error message.
- Constraints: must follow KVKK for user data; must support iOS + Android.

## Design
- External: Payment API (iyzico), Push notification service
- Components: Mobile App, REST API, DB, Payment Gateway, Notifier
- Trade-off: Flutter vs Native — one codebase means faster, cheaper build; slightly less native feel.

## Build & Test
- Tasks: [DB] orders + payments tables · [API] POST /orders + payment call · [UI] cart + pay screen · [Test] payment flow tests · [Docs] document POST /orders
- Git: feature/pre-order-pay → 3 small commits → PR → review → merge
- DoD: tests pass, linter clean, PR approved, no secrets committed
- Tests: unit (price formatter), integration (POST /orders hits DB), E2E (user orders → pays → shop notified)

## Release
- Checklist: migrations run on prod, QA sign-off, smoke test menu → pay → shop notify
- Rollback: disable payment feature flag + revert to previous stable tag
- Changelog: Added pre-order and in-app payment (iyzico)

## Operate
- Maintenance: corrective (payment bugs), adaptive (iyzico API changes)
- Monitor: failed payment rate, 5xx on /orders, order notification delay
- Incident: alert on failed-payment spike → disable flag, rollback, refund customers, post-mortem


# TASK 3 

Model: Hybrid
Why:
- Payment and security are planned upfront (Waterfall) — real money and KVKK risk.
- UI and order flow are built in 2-week sprints (Agile) — fast user feedback.
- Shop onboarding can iterate after the MVP core is stable.


# TASK 4 

1. Phase I underestimate most: Requirements & Design — early on I mixed up the phases (like design vs implementation) and what each one produces.

2. How I'll practice: Re-read my Day 1–4 notes and write a one-line summary of what each phase produces, in my own words, by end of today.

