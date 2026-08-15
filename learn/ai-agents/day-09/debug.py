
def trace_messages(result):
    for i, msg in enumerate(result["messages"]):
        print(f"[{i}] {type(msg).__name__}")
        print(f"    content: {msg.content}")
        if hasattr(msg, "tool_calls") and msg.tool_calls:
            print(f"    tool_calls: {msg.tool_calls}")