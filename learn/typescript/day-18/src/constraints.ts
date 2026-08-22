function logId<T extends {id: string}>(item: T):void{
    console.log(item.id);
}


logId({id: "menu-1", name: "Breakfast Menu"});
logId({id: "menu-66", total: 45});

function describe<T extends { id: string } & { name: string }>(item: T): string {
  return `${item.id}: ${item.name}`;
}
console.log(describe({ id: "drink-1", name: "Latte", price: 45 }));


function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
  return obj[key];
}
const menuItem = { id: "1", name: "Latte", price: 45 };
console.log(getProperty(menuItem, "name"));  
console.log(getProperty(menuItem, "price")); 


function getNameUnsafe<T>(item: T): string {
  return (item as any).name; 
}


function getNameSafe<T extends { name: string }>(item: T): string {
  return item.name; 
}
console.log(getNameSafe({ name: "Espresso", price: 30 }));



logId({ id: "menu-1", name: "Breakfast Menu" });
console.log(describe({ id: "drink-1", name: "Latte", price: 45 }));
console.log(getProperty(menuItem, "name"));
console.log(getNameSafe({ name: "Espresso", price: 30 }));