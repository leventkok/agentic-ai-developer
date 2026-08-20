// Task 2: if narrowing
function printLength(value: string | null): void {
  if (value === null) {
    console.log("No value");
    return;
  }
  console.log("Length:", value.length);
}

function printOrder(drink: string | undefined): void {
  if (drink === undefined) {
    console.log("No drink selected");
  } else {
    console.log(`Order: ${drink}`);
  }
}

// Task 3: switch
type Size = "small" | "medium" | "large";
function priceForSize(size: Size): number {
  switch (size) {
    case "small":
      return 30;
    case "medium":
      return 45;
    case "large":
      return 55;
  }
}

// Task 3: loop
const menu = ["latte", "espresso", "filter"];
for (const item of menu) {
  console.log(item.toUpperCase());
}

// Task 4: null handling
function getDiscount(code: string | null): number {
  if (code === null) return 0;
  return code === "WELCOME" ? 10 : 0;
}

const guestName: string | undefined = undefined;
console.log("Hello,", guestName ?? "Guest");

// Run demos
printLength("latte");
printLength(null);
printOrder("espresso");
console.log("Medium price:", priceForSize("medium"));
console.log("Discount:", getDiscount("WELCOME"));
console.log("Discount (null):", getDiscount(null));
