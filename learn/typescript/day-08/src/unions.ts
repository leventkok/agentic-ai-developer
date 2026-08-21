





type DrinkSize = "small" | "medium" | "large";

function sizeLabel(size: DrinkSize): string{
    return `Size: ${size}`;

}


sizeLabel("medium");



type PriceInput = string | number;

function formatPrice(input: PriceInput): string{
    if(typeof input === "string"){
        return input;
    }
    return `$${input} TL`;
}



type Person = {
    name: string;
}

type Employee = {employeeId: number};

type Staff = Person & Employee;


const barista : Staff = {
    name: "Ali",
    employeeId: 101
};



type MenuItem = {
    name: string;
    price: number;
};

type Discountable = {
    discountPercent: number;
}


type SpecialItem = MenuItem & Discountable;


const mocha: SpecialItem = {
    name: "Mocha",
    price: 50,
    discountPercent: 10

}



type OrderSuccess = {
    kind: "success";
    drink: string;
    price: number;
}

type OrderError = {
    kind: "error";
    message: string;
}

type OrderResponse = OrderSuccess | OrderError;



function handleOrder(response: OrderResponse): void {
    switch (response.kind) {
      case "success":
        console.log(`${response.drink}: ${response.price} TL`);
        break;
      case "error":
        console.log("Error:", response.message);
        break;
    }
  }


const ok: OrderResponse= {
    kind: "success",
    drink: "latte",
    price: 45,};


const fail: OrderResponse = {
    kind: "error",
    message: 'drink not available'
};

handleOrder(ok);
handleOrder(fail);



type Input = string | number;

function printInput(value: Input): void{
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
    if ("address" in order) {
      console.log("Deliver to:", order.address);
    } else {
      console.log("Table:", order.tableNumber);
    }
}


