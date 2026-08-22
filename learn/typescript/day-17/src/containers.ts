
class Stack<T> {
  private items: T[] = [];

  push(item: T): void {
    this.items.push(item);
  }

  pop(): T | undefined {
    return this.items.pop();
  }

  peek(): T | undefined {
    return this.items[this.items.length - 1];
  }
}

interface User {
  id: number;
  name: string;
}

interface KeyValueCache<K, V> {
  get(key: K): V | undefined;
  set(key: K, value: V): void;
}

class InMemoryCache<K, V> implements KeyValueCache<K, V> {
  private store = new Map<K, V>();

  get(key: K): V | undefined {
    return this.store.get(key);
  }

  set(key: K, value: V): void {
    this.store.set(key, value);
  }
}

const orders = new Stack<string>();
orders.push("Latte");
orders.push("Espresso");
console.log("Pop:", orders.pop());
console.log("Peek:", orders.peek());

const userCache = new InMemoryCache<string, User>();
userCache.set("ali", { id: 1, name: "Ali" });
console.log("User:", userCache.get("ali")?.name);

// Task 4: wrong types 
// userCache.set(123, { id: 1, name: "Ali" });
// userCache.set("ali", "not a user");
