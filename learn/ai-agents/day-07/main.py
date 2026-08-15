


from pathlib import Path

from dotenv import load_dotenv
from langchain.agents import create_agent
from langchain_community.tools import DuckDuckGoSearchRun
from langchain_google_genai import ChatGoogleGenerativeAI
from tools.menu import get_menu_price, calculate_order_total
from langgraph.checkpoint.memory import MemorySaver



_day_dir = Path(__file__).resolve().parent
_repo_root = _day_dir.parents[2]

checkpointer = MemorySaver()


def _load_env(path: Path) -> None:
    if not path.exists():
        return
    for encoding in ("utf-8", "utf-16"):
        try:
            load_dotenv(path, encoding=encoding)
            return
        except UnicodeDecodeError:
            continue


_load_env(_day_dir / ".env")
_load_env(_repo_root / ".env")

search = DuckDuckGoSearchRun()
tools = [get_menu_price, calculate_order_total, search]

llm = ChatGoogleGenerativeAI(model="gemini-3.6-flash", temperature=0)
agent = create_agent(
    model=llm,
    tools=tools,
    system_prompt=(
        "You are a RotaCoffee assistant. "
        "Use menu tools for drink prices and order totals. "
        "Use search for general knowledge. "
        "Remember what the user told you in this conversation."
    ),
    checkpointer=checkpointer,
    debug=False,
)


def ask(question: str, config: dict) -> str:
    result = agent.invoke({"messages": [{"role": "user", "content": question}]}, config=config)
    content = result["messages"][-1].content
    if isinstance(content, list):
        return " ".join(
            part.get("text", "") for part in content if isinstance(part, dict)
        )
    return content


if __name__ == "__main__":

    config = {"configurable": {"thread_id": "demo-1"}}

    print(ask("My name is Levent and I like latte.", config))
    print(ask("Whats my favorite drink?", config))
    print(ask("What is my name?", config))

    config2 = {"configurable": {"thread_id": "demo-2"}}
    print(ask("What is my name?", config2))


    
