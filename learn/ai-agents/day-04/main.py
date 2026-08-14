from pathlib import Path

from dotenv import load_dotenv
from langchain.agents import create_agent
from langchain_community.tools import DuckDuckGoSearchRun
from langchain_google_genai import ChatGoogleGenerativeAI



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

search = DuckDuckGoSearchRun()
tools = [search]

llm = ChatGoogleGenerativeAI(model="gemini-3.6-flash", temperature=0)
agent = create_agent(
    model=llm,
    tools=tools,
    system_prompt=(
        "You are a research assistant. Use search to find accurate info. "
        "Be concise and cite sources when possible."
    ),
    debug=True,
)


def ask(question: str) -> str:
    result = agent.invoke({"messages": [{"role": "user", "content": question}]})
    return result["messages"][-1].content


if __name__ == "__main__":
    questions = [
        "What is the ReAct pattern in AI agents?",
        "Who created LangChain?",
    ]
    for q in questions:
        print(f"\n{'=' * 50}")
        print(f"Q: {q}")
        answer = ask(q)
        print(f"A: {answer}")
