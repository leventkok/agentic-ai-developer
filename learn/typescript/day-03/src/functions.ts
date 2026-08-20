


function formatPrice(name: string, price: number): string{
    return `${name}: ${price} TL`
}

console.log(formatPrice("latte", 45));


function greetConsumer(name: string, title?: string): string{
    if(title){
        return `Hello, ${title} ${name}!`;

    }
    return `Hello, ${name}!`;
}


console.log(greetConsumer("John"));
console.log(greetConsumer("John", "Ms."));


function makeCoffee(drink: string, milk: boolean = false, note?: string): string{
    let order = milk ? `${drink} with milk` : drink;

    if(note){
        order = `${order} (${note})`;
    }
    return order;
}


console.log(makeCoffee("latte"));
console.log(makeCoffee("latte", true));
console.log(makeCoffee("espresso", true, "no sugar"));

const applyDiscount = (price: number, percent: number): number => {
    return price - price * (percent / 100);
}


const prices = [45, 35, 30];
const discount = prices.map((p) => applyDiscount(p, 10));

console.log("Discounted:", discount);