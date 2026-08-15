from pathlib import Path
from dotenv import load_dotenv
from agents.barista import create_barista
from agents.researcher import create_researcher


_day_dir = Path(__file__).resolve().parent
_repo_root = _day_dir.parents[2]


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

barista_agent = create_barista()
researcher_agent = create_researcher()

MENU_KEYWORDS = ["how much", "price", "order", "total", "menu", "cost"]

def route(question: str) -> str:
    q = question.lower()
    if any(word in q for word in MENU_KEYWORDS):
        return "barista"
    return "researcher"


def ask(question: str, config: dict) -> str:
    agent_name = route(question)
    agent = barista_agent if agent_name == "barista" else researcher_agent
    result = agent.invoke({"messages": [{"role": "user", "content": question}]}, config=config)
    content = result["messages"][-1].content
    if isinstance(content, list):
        content = " ".join(
            part.get("text", "") for part in content if isinstance(part, dict)
        )
    return f"[{agent_name}] {content}"




if __name__ == "__main__":

    import sys
    from evaluate import run_evaluation


    if len(sys.argv) > 1 and sys.argv[1] == "eval":
        run_evaluation(ask, route)
    else:
        config = {"configurable": {"thread_id": "demo-2"}}
        print(ask("How much is a latte?", config))
        print(ask("Order 2 lattes and 1 espresso — what's the total?", config))

        config2 = {"configurable": {"thread_id": "demo-3"}}
        print(ask("What is the origin of espresso?", config2))
        print(ask("Who invented the latte?", config2))
     
    