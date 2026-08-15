# Day 7 — Memory and State

## Concepts
- **Short-term memory:** conversation history within one session (thread)
- **Long-term memory:** persists across sessions (vector store — Day 7 intro only)
- **thread_id:** LangGraph uses this to separate conversations

## What changed from Day 6
- Added `MemorySaver` checkpointer
- Pass `config={"configurable": {"thread_id": "..."}}` to every `invoke`

## Test results
Thread "user-1":
- Q: "I'm Levent, how much is a latte?" → "A Latte at RotaCoffee is 45 TL..."
- Q: "What's my favorite drink?" → "Your favorite drink is latte!"
- Q: "What's my name?" → "Your name is Levent!"

New thread "user-2":
- Q: "What's my name?" → "You haven't told me your name yet!"

## Observations
- With the same thread_id, the agent remembered my name and my favorite drink 
  across different messages.
- With a new thread_id, the agent had no memory and asked for my name again.
- This shows short-term memory is tied to the thread_id — the MemorySaver keeps 
  history per thread.