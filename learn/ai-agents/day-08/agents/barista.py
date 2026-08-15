from langchain.agents import create_agent
from langchain_google_genai import ChatGoogleGenerativeAI
from tools.menu import get_menu_price, calculate_order_total
from langgraph.checkpoint.memory import MemorySaver



def create_barista():
    """Create a barista agent that can make coffee and tea"""
    llm = ChatGoogleGenerativeAI(model="gemini-3.6-flash", temperature=0)
    tools = [get_menu_price, calculate_order_total]
    return create_agent(
        model = llm,
        tools = tools,
        system_prompt = (
            "You are the RotaCoffee barista. "
            "Help with menu prices and order totals. "
            "Use menu tools — do not search the web."
        ),
        checkpointer = MemorySaver(),
    )