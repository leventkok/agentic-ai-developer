


//type DrinkName = "latte" | "espresso" | "filter";

const prices: Record<DrinkName, number> = {
    latte: 45,
    espresso: 35,
}

function getPrice(name: DrinkName): number{
    return prices[name];
}

console.log(getPrice("latte"));



const config = {
    shop: "MasterFabric Cafe",
    maxOrders: 100,

}as const;


const ROUTES = {
    menu: "/menu",
    orders: "/orders",
    checkout: "/checkout",
}as const;


type RouteKey = keyof typeof ROUTES;
type RoutePath = (typeof ROUTES)[RouteKey];



function navigate(key: RouteKey): void{
    console.log(`Navigating to: ${ROUTES[key]}`);
}


navigate("menu");


//type OrderStatus = "pending" | "ready" | "canceled";


function printStaus(status: OrderStatus): void{
    console.log(status);
}




const SHOP_CONFIG = {
    name: "MasterFabric Cafe",
    currency: "TRY",
    roles: ["admin", "barista", "guest"] as const,
    drinks: {
      latte: { price: 45, sizes: ["small", "medium", "large"] as const },
      espresso: { price: 35, sizes: ["small", "medium"] as const },
    },
    orderStatuses: ["pending", "ready", "cancelled"] as const,
  } as const;
  
  type Role = (typeof SHOP_CONFIG.roles)[number];
  
  type DrinkName = keyof typeof SHOP_CONFIG.drinks;
  
  type OrderStatus = (typeof SHOP_CONFIG.orderStatuses)[number];
  
  function getDrinkPrice(drink: DrinkName): number {
    return SHOP_CONFIG.drinks[drink].price;
  }
  
  function canAccess(role: Role, required: Role): boolean {
    if (role === "admin") return true;
    if (role === "barista" && required !== "admin") return true;
    return role === required;
  }
  
  console.log(getDrinkPrice("latte"));
  console.log(canAccess("barista", "guest"));