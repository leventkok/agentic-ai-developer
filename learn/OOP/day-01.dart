

// TASK 1
// A section


Map<String, int> menu = {
  "latte": 10,
  "espresso": 5,
  "filter": 7,
};

List<String> order = [];  

void addOrder(String item) {
  order.add(item);
}

int totalOrder() {
  int total = 0;
  for (String item in order) {
    total += menu[item] ?? 0;
  }
  return total;
}

// B section

class OrderCoffee{
    String coffeName;
    int coffePrice;

    OrderCoffee(this.coffeeName, this.coffeePrice);

    void addCoffe(String coffeeName, int coffeePrice){
        this.coffeeName = coffeeName;
        this.coffeePrice = coffeePrice;
    }

    int totalOrder(){
        return this.coffeePrice;
    }
}


// TASK 2 

/*
CoffeeShop = holds menu items and receives orders
Order = track items and total price
Customer = can add items to order and payments
MenuItem = it can holds name, prices, sizes, and the order it belongs to

*/


// TASK 3

/*
Class = A template for creating objects = Example = Order
Object = An instance of a class = Example = Order1, Order2...
Attributes = Properties of an object = Example = Order.items, Order.totalPrice
Methods = Functions of an object = Example = Order.addItem(), Order.removeItem()

*/

// TASK 4

// DART