

// strings

const drinkName: string = "latte";
const shopName: string = "MasterFabric Cafe";


// numbers
const price: number = 45;
const temperature: number = 42;


// booleans

const isAvailable: boolean = true;
const isDecaf: boolean = false;

console.log(`${drinkName} costs: ${price} TL`);



// explicit types
let cupsSold: number = 0;
cupsSold = cupsSold + 1;

//infered types
const menuItem = "espresso";
const menuPrice = 35;

console.log(`Sold ${cupsSold} cups. Item: ${menuItem}, price: ${menuPrice} TL`);


// task 3 unkown (safe)

function handleOrder(raw: unknown): void{
    if(typeof raw === "string"){
        console.log("Drink order:", raw);
    }else if(typeof raw == "number"){
        console.log("Quantity:", raw);
    }else{
        console.log("Invalid order - expected string or number");
    }
}


handleOrder("latte");
handleOrder(10);
handleOrder(true);


// assertions


const apiData: unknown = "filter";

if(typeof apiData == "string"){
    console.log("Narrowed:", apiData.toUpperCase());
}


const asserted = apiData as string;

console.log("Asserted:", asserted.toUpperCase());