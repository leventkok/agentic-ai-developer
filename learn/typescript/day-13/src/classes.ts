// Day 13: Classes and Access Modifiers


class Product{
    constructor(public name: string, public price: number, private cost: number) {}
    getMargin(): number{
        return this.price - this.cost;
    }

    describe(): string{
        return `${this.name} - ${this.price} TL - ${this.getMargin()} TL margin`;
    }
}


class Beverage extends Product{
    constructor(
        name: string,
        price: number,
        cost: number,
        protected category: "hot" | "cold"
    ){
        super(name, price, cost);
    }

    categoryLabel(): string{
        return `${this.category} drink`
    }
}


const latte = new Beverage("Latte", 45, 30, "hot");
// latte.cost  => error
// latte.category => error
console.log(latte.describe());
console.log(latte.categoryLabel());




class Order{
    constructor(public readonly id: number, public item: string, public total: number) {}
}

type MenuEntry = {
    name: string;
    price: number;
};


const plainItem: MenuEntry ={
    name: "Espresso",
    price: 30
};

class MenuEntryClass{
    constructor(
        public name: string,
        public price: number,
    ){}

    discount(percent: number): number{
        return this.price * (1 - percent / 100);
    }

}

const classItem = new  MenuEntryClass("Espresso", 30);
console.log("Discounted: ", classItem.discount(10));



// Use plain objects for data; use classes when you need methods + encapsulation.
const product = new Product("Croissant", 35, 12);
console.log(product.describe());

const order = new Order(101, "Latte", 45);
console.log(`Order #${order.id}: ${order.item}`);

console.log("Plain:", plainItem.name);
console.log("Class discount:", classItem.discount(10));
