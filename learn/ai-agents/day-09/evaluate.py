from pathlib import Path
from dotenv import load_dotenv
from agents.barista import create_barista
from agents.researcher import create_researcher


EVAL_CASES = [
    {
        "question": "How much is a latte?",
        "expected_agent": "barista",
        "expected_in_response": ["45"]
    },
    {
        "question": "Order 2 lattes and 1 espresso — what's the total?",
        "expected_agent": "barista",
        "expected_in_response": ["125"],
    },
    {
        "question": "What is the origin of espresso?",
        "expected_agent": "researcher",
        "expected_in_response": [],
    },
    {
        "question": "Who invented the latte?",
        "expected_agent": "researcher",
        "expected_in_response": [],
    },
    {
        "question": "Tell me about latte art history",
        "expected_agent": "researcher",
        "expected_in_response": [],
    }
]



def score_case(question, expected_agent, expected_in_response, ask_fn, route_fn, config):
    actual_agent = route_fn(question)
    response = ask_fn(question, config)
    routing_ok = actual_agent == expected_agent
    content_ok = all(s.lower() in response.lower() for s in expected_in_response)

    return routing_ok, content_ok, response



def run_evaluation(ask_fn, route_fn):
    config = {"configurable": {"thread_id": "eval-run"}}

    routing_pass = 0
    content_pass = 0
    total = len(EVAL_CASES)

    print("=== Evaluation REPORT ===")

    for case in EVAL_CASES:
        routing_ok, content_ok, response = score_case(
            case["question"],
            case["expected_agent"],
            case["expected_in_response"],
            ask_fn,
            route_fn,
            config
        )


        if routing_ok:
            routing_pass += 1
        if content_ok:
            content_pass += 1
        
        if routing_ok and content_ok:
            status = "PASS"
        elif routing_ok:
            status = "ROUTING_PASS"
        else: 
            status = "FAIL"

        print(f"{status} | {case['question']}")

        if not routing_ok:
            print(f"expected agent: {case['expected_agent']}, got: {route_fn(case['question'])}")
        if not content_ok and case["expected_in_response"]:
            print(f"response missing: {case['expected_in_response']}")
            print(f" got: {response[:80]}...")
    
    print(f"\nRouting: {routing_pass}/{total}")
    print(f"Content: {content_pass}/{total}")