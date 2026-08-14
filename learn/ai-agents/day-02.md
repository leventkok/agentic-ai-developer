
# ELI5
Q: How do LLMs act as the "brain" of an AI agent?
A: The LLM is the brain — it reads context and decides what to do. It does not act directly; it chooses tools or writes a reply.

# Scenario
Q: Explain tool use with a concrete example (not weather — pick your own: GitHub, shopping, coffee orders, etc.). Show: user request → LLM decision → tool call → result → final answer.

A:
- The user enters the prompt: "I need a sentiment analysis of the reviews for app X."
- Observe -> LLM read the user prompt
- Think -> I need the reviews app X 
- Act -> Tool call
- Act -> Fetchs store reviews for app X 
- Observe -> Sees reviews and analyze the reviews
- Think -> Task complete
- Act -> replies to user
-Final answer: "App X sentiment: 72% positive, main complaint: slow loading."


# Deep Dive
Q: Why is memory important? Explain short-term vs long-term memory with one example each.
A: important for keeping the conversation going, For example, in a short-term conversation, it does not remember the first output after five outputs. However, in long-term conversations, it remembers details—such as user preferences—just as it did on the very first day.

# Trace the LOOP

Q: Pick a real task (something you did with Cursor or ChatGPT). Write 4–6 steps labeled Observe / Think / Act.

A:
Task: Fix the "go run" error in my Go calculator

1. Observe: I ran `go run . 10 + 5` and got a compile error about the calc package.
2. Think: The error says "expected package, found EOF" — the calculator.go file might be empty.
3. Act: I open calculator.go and add the package declaration and Calculate function.
4. Observe: I run it again and now get different errors about `reulst` and `fmt.println`.
5. Think: These are typos — wrong variable name and wrong capitalization.
6. Act: I fix the typos and run again — now it prints "10 + 5 = 15".


# Connection
Q: Which is more dangerous in production — a bad LLM answer or a bad tool call? Why?
A: Bad tool call more dangerous than bad llm answer. Because tools can act something.