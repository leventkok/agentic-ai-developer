

type Input = string | number;

function printInput(value: Input): void {
    if(typeof value === "string"){
        console.log(value.toUpperCase());
    }else{
        console.log(value.toFixed(2));
    }
}



type DeliveryOrder = {type: "delivery"; address: string};
type PickupOrder = {type: "pickup"; tableNumber: number};
type Order = DeliveryOrder | PickupOrder;

function printOrder(order: Order): void {
    if("address" in order){
        console.log("Deliver to:", order.address);
    }else {
        console.log("Table:", order.tableNumber);
    }
}


type User = {
    id: number;
    name: string;
}

function isUser(value: unknown): value is User{
    return (
        typeof value === "object" &&
        value !== null &&
        "id" in value &&
        "name" in value &&
        typeof (value as User).id === "number" &&
        typeof (value as User).name === "string"
    );
}


function processInput(input: unknown): void{
    if(isUser(input)){
        console.log(input.name);
    }else{
        console.log("Invalid input");
    }
}

type MenuItem = {
    name: string; price: number
};

function isMenuItem(value: unknown): value is MenuItem{
    return(
        typeof value === "object" &&
        value !== null &&
        "name" in value &&
        "price" in value &&
        typeof (value as MenuItem).name === "string" &&
        typeof (value as MenuItem).price === "number"
    );
}


type OrderResponse = | {kind: "success"; drink: string; price: number} | {kind: "error"; message: string};

function handleOrder(response: OrderResponse): string{
    switch(response.kind){
        case "success":
            return `${response.drink}: ${response.price} TL`;
        case "error":
            return `Error: ${response.message}`;
        default:
            const _exhaustive: never = response;
            return _exhaustive;             
    }
}








