from langchain.agents import create_agent
from langchain_community.tools import DuckDuckGoSearchRun
from langchain_google_genai import ChatGoogleGenerativeAI
from langgraph.checkpoint.memory import MemorySaver

def create_researcher():
    """Create a researcher agent that can search the web for information"""
    llm = ChatGoogleGenerativeAI(model="gemini-3.6-flash", temperature=0)
    tools = [DuckDuckGoSearchRun()]
    return create_agent(
        model = llm,
        tools = tools,
        system_prompt = (
            "You are a coffee research assistant for RotaCoffee. "
            "Use search for general knowledge, history, and facts. "
            "Do not make up menu prices — you don't have the menu."
        ),
        checkpointer = MemorySaver(),
    )