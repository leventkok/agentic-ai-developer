/*
 DAY 5 - RotaCoffee checkout capstone

    MenuItem: one product with a name and price, validates its own data
    Order: holds a list of items, calculates total, tracks paid status
    Customer: has a balance, decides if it can pay and pays an order

*/
// TASK 1


class MenuItem{
    final String name;
    final int price;
    MenuItem(this.name, this.price){
        if(name.isEmpty) throw Exception('Name cannot be empty');
        if(price < 0) throw Exception('Price cannot be negative');
    }

    String describe() => 'MenuItem(name: $name, price: $price)';
}

class Order{
    final List<MenuItem> items = [];
    bool isPaid = false;

    Order();


    void addItem(MenuItem item){
        items.add(item);
    }

    int totalPrice() => items.fold(0, (sum, item) => sum + item.price);
    bool isEmpty() => items.isEmpty;
    void markAsPaid(){
        isPaid = true;
    }


}


class Customer{
    int balance;
    Customer(this.balance){
        if(balance < 0) throw Exception('Balance cannot be negative');
    }

    void pay(Order order){
        if(order.isEmpty()) return;
        if(balance < order.totalPrice()) return;
        balance -= order.totalPrice();
        order.markAsPaid();
    }
    bool canPay(Order order) => balance >= order.totalPrice();


}


void main() {

  print('=== Scenario 1: Happy path ===');
  final order1 = Order();
  order1.addItem(MenuItem('Latte', 45));
  order1.addItem(MenuItem('Espresso', 35));
  final customer1 = Customer(100);
  customer1.pay(order1);
  print('Paid: ${order1.isPaid} | Balance left: ${customer1.balance}');


  print('=== Scenario 2: Insufficient funds ===');
  final order2 = Order();
  order2.addItem(MenuItem('Latte', 45));
  final customer2 = Customer(30);        
  customer2.pay(order2);
  print('Paid: ${order2.isPaid} | Balance: ${customer2.balance}');


  print('=== Scenario 3: Empty order ===');
  final order3 = Order();               
  final customer3 = Customer(100);
  customer3.pay(order3);
  print('Paid: ${order3.isPaid} | Balance: ${customer3.balance}');

  print('=== Scenario 4: Invalid item ===');
  try {
    final bad = MenuItem('', -5);        
    print(bad.describe());
  } catch (e) {
    print('Error: $e');
  }
}
