// Day 14: Structural Typing and Declaration Merging
// Work through Tasks 1, 2, and 4 here. Task 3 lives in shop.d.ts
// Then run: npm run build && npm start

import { initShop } from "./shop";



// TypeScript is structural: compatibility depends on shape, not type names.
// In nominal languages (Java), ali and Waiter would be incompatible unless ali explicitly implements Waiter.

interface Waiter {
    name: string;
    serve(table: number): void;
}


const ali = {
    name: "Ali",
    serve(table: number){
        console.log(`Serving table ${table}`);
    }
};


function assignWaiter(w: Waiter): void{
    w.serve(5);
}

assignWaiter(ali);


interface Dimensions{
    width: number;
}


interface Dimensions{
    height: number;
}


const tray: Dimensions = {width: 30, height: 20};

console.log("Tray:", tray.width, "x", tray.height);



const config = {
    shopName: "Masterfabric Cafe",
    version: "1.0.0",
};

initShop(config);