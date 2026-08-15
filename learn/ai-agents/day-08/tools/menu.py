from langchain.tools import tool

@tool
def get_menu_price(drink: str) -> str:
    """Look up the price of a drink on the menu, Input: drink name -> drink price like "latte" → "Latte: 45 TL" """
    menu = {"latte" : 45, "espresso" : 35, "filter" : 30}
    price = menu.get(drink.lower())
    if price is None:
        return f"Sorry, {drink} is not on the menu."
    return f"{drink.title()}: {price} TL"

@tool
def calculate_order_total(items: list[str]) -> str:
    """Calculate the total price of a list of items, Input: list of item names -> total price like "latte, espresso, latte" → "Total: 125 TL" """
    menu = {"latte" : 45, "espresso" : 35, "filter" : 30}
    total = 0
    for item in items:
        price = menu.get(item.lower())
        if price is not None:
            total += price
    return f"Total: {total} TL"