// TASK 1

class MenuItem {
  final String name;
  final int price;
  MenuItem(this.name, this.price);
}

class Order {
  final List<MenuItem> items = [];
  bool isPaid = false;

  void addItem(MenuItem item) => items.add(item);
  int totalPrice() => items.fold(0, (s, i) => s + i.price);
  bool isEmpty() => items.isEmpty;
  void markAsPaid() => isPaid = true;
}

// TASK 2

class Customer {
  int balance;
  Customer(this.balance);

  bool canPay(Order order) => balance >= order.totalPrice();

  void pay(Order order) {
    if (order.isEmpty() || !canPay(order)) return;
    balance -= order.totalPrice();
    order.markAsPaid();
  }
}

// TASK 4

void main() {
  final latte = MenuItem('Latte', 45);
  final espresso = MenuItem('Espresso', 35);

  final order = Order();
  order.addItem(latte);
  order.addItem(espresso);

  final customer = Customer(100);
  customer.pay(order);

  print('Paid: ${order.isPaid}');
  print('Balance: ${customer.balance}');
  print('Total: ${order.totalPrice()}');
}
