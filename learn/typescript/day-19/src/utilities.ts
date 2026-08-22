// Day 19: keyof, Defaults, and Reusable Abstractions
// Work through the 4 tasks below, then run: npm run build && npm start


function pick<T, K extends keyof T>(obj: T, keys: K[]): Pick<T, K> {
    const result = {} as Pick<T, K>;
    for (const key of keys){
        result[key] = obj[key];
    }
    return result;
}



type MenuItem = {
    id: string;
    name: string;
    price: number;
    category: string;
};

const latte: MenuItem = {
    id: "1",
    name: "Latte",
    price: 45,
    category: "Hot",

};

const publicView = pick(latte, ["name", "price"]);




function wrapInArray<T = string>(value: T): T[] {
    return [value];
  }
  const strings = wrapInArray("Latte");     
  const numbers = wrapInArray<number>(45); 



  function partial<T>(obj: T): Partial<T> {
    return { ...obj };
  }
 
  const picked = pick(latte, ["name", "price"]);
  const update = partial(picked);
  update.price = 50; 
  console.log("Update:", update);


  // Library example: Array<T> — T defaults to the element type you pass.
// Promise<T> — T is the resolved value type inferred from async functions.
// Reading these signatures shows generics are everywhere in npm packages.

console.log("Public view:", publicView);
console.log("Wrapped:", wrapInArray("Espresso"));
console.log("Update:", partial(pick(latte, ["name", "price"])));