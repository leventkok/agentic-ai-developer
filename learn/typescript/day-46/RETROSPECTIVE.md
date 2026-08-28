# Retrospective — Day 50 (Node.js TypeScript capstone)

## What went well
- Clear folder layers: config → storage → service → handler → server
- Typed env validation fails fast at boot
- Result + discriminated AppError for predictable API responses
- Integration test hits real HTTP server with temp file storage

## What was hard
- Separating server creation from listen for testability
- Parsing JSON bodies from IncomingMessage streams
- Node ESM requires `.js` extensions in imports

## Patterns to reuse
- **loadConfig()** — never read `process.env` in handlers
- **FileNoteStore** — fs/promises with ENOENT handling
- **errorResponse()** — map domain errors to HTTP status
- **Thin handlers** — parse, delegate, respond
