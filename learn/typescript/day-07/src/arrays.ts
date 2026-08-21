type MenuItem = {
    name: string;
    price: number;
  };
  // Task 1: typed arrays
  const drinkNames: string[] = ["latte", "espresso", "filter"];
  const menu: MenuItem[] = [
    { name: "latte", price: 55 },
    { name: "espresso", price: 35 },
  ];
  console.log(menu[0].name);
  // Task 2: tuples
  type OrderResult = [drink: string, price: number, success: boolean];
  const result: OrderResult = ["latte", 55, true];
  const [drink, itemPrice, ok] = result;
  console.log(drink, itemPrice, ok);
  type Result = [string | null, number];
  function getPrice(name: string): Result {
    if (name === "latte") return [null, 45];
    return ["Drink not found", 0];
  }
  const [error, orderPrice] = getPrice("latte");
  if (error === null) {
    console.log("Price:", orderPrice);
  }
  // Task 3: readonly array
  const readonlyMenu: readonly string[] = ["latte", "espresso"];
  console.log("Readonly menu:", readonlyMenu);
  // Task 4: iterate with type safety
  for (const name of drinkNames) {
    console.log(name.toUpperCase());
  }
  const total = menu.map((item) => item.price).reduce((sum, p) => sum + p, 0);
  console.log("Total:", total);
  