# TASK 1 

# RotaCoffee — Mini Spec (MVP Pre-Order)

## Problem
People waste time waiting in long queues for their morning coffee, and when they're
in a hurry they often skip it. Local coffee shops also struggle to compete with big
chains and to reach their customers with promotions.

## Users
- Coffee drinker (consumer)
- Barista / shop owner
- Admin

## Delivery model
Hybrid — we plan the payment and user-data parts upfront (Waterfall) because they
carry risk, but build the UI and features in short Agile sprints for fast feedback.

## User stories (5)
1. As a coffee drinker, I want to pre-order my coffee before I arrive, so that I can skip the morning queue.
2. As a coffee drinker, I want to know when my coffee is ready, so that I don't wait around.
3. As a barista, I want to manage orders through an online system during busy hours, so that I don't mix up customers.
4. As a coffee shop owner, I want to send promotions to customers through the app, so that I can compete with big chains.
5. As an admin, I want to add new local coffee shops to the app, so that we can grow our market share.

## Acceptance criteria (core flow — pre-order + barista)
Pre-order (coffee drinker):
Given I am logged in and viewing a shop menu
When I select a latte and tap "Order"
Then I see an order confirmation with estimated pickup time

Order management (barista):
- Barista can see new orders in a list
- Order shows customer name, drink, and size
- Barista can mark an order as "Ready"
- Shop receives order within 30 seconds

## Constraints
- Business: Revenue comes from partnerships with coffee shops; many competitors in the market.
- Time: Must launch before the busy season and handle seasonal rush hours.
- Legal: Must follow KVKK for user data and comply with payment regulations.
- Technical: Requires internet connection; must support both iOS and Android.

## MVP scope (Priority 1–3 only)
1. View menu
2. Add to cart
3. Pay + shop receives order
Out of scope for MVP: favorites, scheduled pickup, live tracking, promotions.

## Architecture trade-off
Decision: Flutter | Alternative: Native (Swift + Kotlin) | Why: One codebase means
faster, cheaper development; the trade-off is a slightly less native feel, but it's
worth it for the MVP.


# TASK 2


## Context diagram

                        ┌─────────────┐
                        │ Payment API │
                        └─────────────┘
                              ▲
                              │ process payment
   ┌──────────┐   order    ┌──┴─────────┐   send notification   ┌──────────────────┐
   │  Coffee  │───────────►│ RotaCoffee │──────────────────────►│ Push Notification│
   │  Drinker │◄───────────│   System   │                       │     Service      │
   └──────────┘  confirm   └──┬──────┬──┘                       └──────────────────┘
                              │      │
                 manage orders│      │ onboard shops
                              ▼      ▼
                     ┌──────────┐  ┌──────────┐
                     │  Barista │  │  Admin   │
                     └──────────┘  └──────────┘


## Component sketch

   ┌──────────────┐        ┌───────────────┐
   │  Mobile App  │        │ Shop Dashboard│
   │ menu,cart,pay│        │ orders, Ready │
   └──────┬───────┘        └───────┬───────┘
          └───────────┬────────────┘
                      ▼
              ┌───────────────┐
              │  REST API     │
              │ auth, orders, │
              │ payment logic │
              └───────┬───────┘
        ┌─────────────┼─────────────┐
        ▼             ▼             ▼
   ┌─────────┐  ┌──────────┐  ┌──────────┐
   │ Database│  │ Payment  │  │ Notifier │
   │ users,  │  │ Gateway  │  │ "ready"  │
   │ orders  │  │          │  │ push     │
   └─────────┘  └──────────┘  └──────────┘



# TASK 3 

1. Payment flow unclear → clarified: Mobile App → REST API → Payment Gateway; order sent to shop only after payment succeeds
2. Pre-order criteria not testable → added confirmation appears within 3 second.
3. MVP scope vague on promotions  → moved promotions explicitly to out of scope