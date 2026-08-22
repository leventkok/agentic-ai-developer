// Day 20: Generics — Practice

type CafeEvents = {
  orderPlaced: { item: string; total: number };
  stockLow: { item: string; remaining: number };
};

class EventEmitter<Events extends Record<string, unknown>> {
  private listeners: {
    [K in keyof Events]?: Array<(payload: Events[K]) => void>;
  } = {};

  on<K extends keyof Events>(
    event: K,
    handler: (payload: Events[K]) => void
  ): void {
    if (!this.listeners[event]) {
      this.listeners[event] = [];
    }
    this.listeners[event]!.push(handler);
  }

  emit<K extends keyof Events>(event: K, payload: Events[K]): void {
    this.listeners[event]?.forEach((handler) => handler(payload));
  }
}

function parseJson<T>(
  text: string,
  guard: (value: unknown) => value is T
): T {
  const parsed: unknown = JSON.parse(text);
  if (!guard(parsed)) {
    throw new Error("Invalid JSON shape");
  }
  return parsed;
}

type MenuItem = { name: string; price: number };

function isMenuItem(value: unknown): value is MenuItem {
  return (
    typeof value === "object" &&
    value !== null &&
    "name" in value &&
    typeof (value as MenuItem).name === "string" &&
    "price" in value &&
    typeof (value as MenuItem).price === "number"
  );
}

function pipe<A, B, C>(
  value: A,
  fn1: (input: A) => B,
  fn2: (input: B) => C
): C {
  return fn2(fn1(value));
}

// Generics prevent bugs in:
// 1. API clients — fetch<User>(...) returns typed data, not any
// 2. State management — useState<CartItem[]>([]) keeps array element types
// 3. Repositories — Repository<T> ensures save/findById use the same entity type

// Run demos
const cafe = new EventEmitter<CafeEvents>();

cafe.on("orderPlaced", (payload) => {
  console.log(`Order: ${payload.item} — ${payload.total} TL`);
});

cafe.on("stockLow", (p) => {
  console.log(`Low stock: ${p.item} (${p.remaining})`);
});

cafe.emit("orderPlaced", { item: "Latte", total: 45 });
cafe.emit("stockLow", { item: "Milk", remaining: 2 });

console.log(
  "Parsed:",
  parseJson('{"name":"Espresso","price":30}', isMenuItem).name
);

const piped = pipe(
  "latte",
  (s) => s.toUpperCase(),
  (s) => ({ name: s, price: 45 })
);
console.log("Piped:", piped.name, piped.price);
