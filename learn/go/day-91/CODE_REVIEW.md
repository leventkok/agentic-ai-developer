# Code Review Guide — Day 91

## Trunk-based workflow

1. Create a **short-lived branch** from `master`
2. Keep PRs **small** (< 400 lines when possible)
3. Merge frequently; avoid long-lived feature branches

## PR description checklist

- [ ] **Why** — business or technical motivation
- [ ] **What** — summary of changes
- [ ] **Test plan** — commands run or manual steps
- [ ] **Risks** — breaking changes, migrations, rollback

## Reviewer checklist

| Area | Look for |
|------|----------|
| **Correctness** | Edge cases, error paths, idempotency |
| **Tests** | New behavior covered; no flaky sleeps |
| **Security** | Auth on protected routes; no secrets in repo |
| **Readability** | Clear names; domain rules in domain layer |
| **Performance** | No N+1 queries; cache invalidation on writes |

## Giving feedback

- Comment on **code**, not people
- Label optional style notes as **nit**
- Suggest alternatives when requesting changes
- Approve when blockers are resolved

## Commands before opening a PR

```powershell
go test ./...
./scripts/verify-quality.ps1
```
