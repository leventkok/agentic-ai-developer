// Day 16: Generic Functions
// Work through the 4 tasks below, then run: npm run build && npm start




function identity<T>(value: T): T{
    return value;
}


const s = identity("latte");
const n = identity(42);
const item = identity({name: "Latte", price: 45});




function first<T>(items: T[]): T | undefined{
    return items[0];

}

function last<T>(items: T[]): T | undefined{
    return items[items.length - 1];
}

function unique<T>(items: T[]): T[]{
    return [...new Set(items)];
}


const drinks = ["latte", "espresso", "latte"];

console.log(first(drinks));
console.log(unique(drinks));



const prices = [45, 30, 55];
console.log(first(prices));


const empty: string[] = [];
const firstEmpty = first<string>(empty);

const guessed = first<string>([]);
console.log("Explicit:", firstEmpty, guessed);


console.log("Identity:", identity("Latte"), identity(45));
console.log("First drink:", first(drinks));
console.log("Unique:", unique(drinks));
console.log("Last price:", last(prices));
console.log("Explicit empty:", first<string>([]));





