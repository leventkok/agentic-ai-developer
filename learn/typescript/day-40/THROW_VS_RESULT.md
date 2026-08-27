# Throw vs Result

| Use Result when | Use throw when |
|-----------------|----------------|
| Validation / not-found (expected) | Out of memory, bug, impossible state |
| Caller must handle every failure | Failure is truly exceptional |
| Public library API (like Day 32 TaskStore) | Internal code, framework boundaries |

This project uses **Result** for user-facing fetch (`fetchUserWithResult`).
**Throw** is fine for unexpected failures (`fetchUserOrThrow` + try/catch).
