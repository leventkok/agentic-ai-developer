# ReAct
Q: Look at your Day 4 debug output for one question. Label each step as Thought, Action, or Observation. (Paste or describe the steps.)
A: 
Thought: I need to find what the ReAct pattern is. I should search.
Action: DuckDuckGoSearchRun("ReAct pattern AI agents")
Observation: ReAct combines reasoning and acting, where the model...
Thought: Now I have enough info to answer.
Action: (final answer)

# ELI5 for a project manager 
Q: Explain ReAct to a non-technical PM in 3–4 sentences. What are the key benefits?
A: ReAct:  looks up real data, then answers. That means fewer wrong answers and answers you can trust for decisions.

# MRKL design
Q: Design an MRKL agent for RotaCoffee that answers: "If we sell 20 lattes and 10 espressos with 15% student discount, what's the total?" Name which module handles: prices, math, and final reply.
A:

- Knowledge: Step1: Retrieve facts -> need search / need DB
- Reasoning: Step2: calculate total -> need math-> Step 3: apply discount
- Language: Step4: format the answer 


# Self-correction scenario 
Q: Your Day 4 agent returns "LangChain was created by Google." How would a self-correction mechanism catch and fix this?
A: 
Thought:  Draft answer — "LangChain was created by Google"
Action:   search("who created LangChain")   ← verify step
Observation: Results say Harrison Chase
Thought:  My answer was wrong — correcting
Action:   respond with "Harrison Chase"

# Architecture pick
Q: For each use case, pick ReAct, MRKL, or Simple chain and justify in one line:
A: 
a) Chatbot that summarizes uploaded PDFs => Simple-chain Architecture => it can summarize PDFs
b) Agent that fixes CI failures via GitHub + terminal => ReAct => it can API calling and it can cath some issues and fix bug
c) Calculator that converts currencies using live rates => MRKL => it can more complex build, it have multi domain


# Phase 1 reflection
Q: In 3 bullets: what was the most important thing you learned in Days 1–5?
A:
- I learned agents types
- Which can i use specialist task
- I learned "how do agents work basicly" 

