# Comparison
Q: What are the fundamental differences in design philosophy between LangChain and LlamaIndex?
A: LangChain usually use agent orchestration. It is like swiss army knife. LlamaIndex is reading and query data. It is like a librarian + brain.  

# Use Case 
Q: You want an agent that chats with PDF documents (e.g. SDLC notes). Which framework fits best and why?
A: Definitely LlamaIndex. Because, it can read our document and answer my question from my data


# Secon Use Case
Q: You want an agent that calls GitHub API + runs terminal commands to fix CI failures. Which framework fits best and why?
A: If i run the terminal command, i can use Cursor. Because, it is coding agent with tools 


# Framework Map
Q: Fill in this table with your own one-line description for each:
A: 
Framework	Best for	                    One thing it handles for you
LangChain   API's the agent can call        general tasks, system prompts and API's agent call 
LlamaIndex  read and query data             reading my homework PDF and answer my questions
AutoGen     Multi- agent conservation       My agent can conservation another agents
CrewAI      Role based teams                My agents have a person like software developor, researcher, exc...



# Personal pick
Q: For your agentic-ai-developer repo journey, which framework do you expect to use first and why?
A: Im expecting use the LangChain because its a general agent type