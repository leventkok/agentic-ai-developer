

function fetchMenuItem(id: number){
    return {id, name: "Latte", price: 45};
}


type MenuItem = ReturnType<typeof fetchMenuItem>;




function printItem(item: MenuItem){
    console.log(`${item.name} - $${item.price}`);
}

function placeOrder(itemName: string, qty: number) {
    return { item: itemName, qty, total: qty * 45 };
}


type PlaeOrderArgs = Parameters<typeof placeOrder>;


function logAndPlace(...args: PlaeOrderArgs){
    console.log("Placing order with:", args);
    return placeOrder(...args);
}



async function fetchUser(id: number): Promise<{id: number, name: string}> {
    return {id, name: "Ali"};
}


async function withRetry<T>(fn: () => Promise<T>): Promise<T> {
    try {
      return await fn();
    } catch {
      return await fn();
    }
}


async function getUserSafe(id: number){
    return withRetry(() => fetchUser(id));
}


type UserResult = ReturnType<typeof getUserSafe>;


function parseInput(input: string): number;
function parseInput(input: number): string;
function parseInput(input: string | number): string | number {
  return typeof input === "string" ? parseInt(input, 10) : String(input);
}

type ParseResult = ReturnType<typeof parseInput>;



const item = fetchMenuItem(1);
printItem(item);

console.log("Order:", logAndPlace("Latte", 2));

getUserSafe(1).then((user) => console.log("User:", user.name));

console.log("Parse number:", parseInput("42"));
console.log("Parse string:", parseInput(99));