

type MenuItem = {
    name: string;
    price: number;
    available: boolean;
};


const latte: MenuItem = {
    name: "Latte",
    price: 3.99,
    available: true,
}


console.log(`${latte.name} costs $${latte.price} TL	`);


type Order = {
    drink: string;
    note?: string;
};

const order1: Order = {drink: "Latte"};
const order2: Order = {drink: "Espresso", note: "Extra shot"};


type Config = {
    readonly shopName: string;
    maxOrders: number;
}

const config: Config = {
    shopName: "Coffee Shop",
    maxOrders: 100,
}

config.maxOrders = 200;

type PriceMenu = {
    [drinkName: string]: number;
}

const menu: PriceMenu = {
    latte: 45,
    espresso: 35,
    filter: 30,
}


console.log(menu.latte);
console.log(menu["espresso"]);
menu.filter = 60;


type Address = {
    city: string;
    street: string;
}

type Shop = {
    name: string;
    address: Address;
    menu: PriceMenu;
};

const shop: Shop = {
    name: "Coffee Shop",
    address: {
        city: "New York",
        street: "123 Main St",
    },
    menu: {
        latte: 45,
        espresso: 35,
    },
};

console.log(shop.address.city);
console.log(shop.menu.latte);
