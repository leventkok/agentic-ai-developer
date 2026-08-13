/*
 DAY 7 — Invariants
 MenuItem invariants:
   1. name is never empty
   2. price is never negative
   3. after discount, price stays >= 0
 Order invariants:
   1. cannot add items to a paid order
   2. totalPrice is never negative
   3. isPaid is only true after successful payment
 Customer invariants:
   1. balance is never negative
   2. pay never reduces balance below 0
   3. deposit only accepts positive amounts
*/

class MenuItem {
  final String name;
  int _price;

  int get price => _price;

  MenuItem(this.name, this._price) {
    if (name.isEmpty) throw ArgumentError('Name cannot be empty');
    if (_price < 0) throw ArgumentError('Price cannot be negative');
  }

  void applyDiscount(int percent) {
    if (percent < 0 || percent > 50) {
      throw ArgumentError('Discount percentage must be between 0 and 50');
    }
    _price = _price - (_price * percent ~/ 100); // integer math keeps price >= 0
  }

  String describe() => 'MenuItem(name: $name, price: $price)';
}

class Order {
  final List<MenuItem> _items = [];
  bool _isPaid = false;

  bool get isPaid => _isPaid;
  List<MenuItem> get items => List.unmodifiable(_items);

  Order();

  void addItem(MenuItem item) {
    if (_isPaid) throw StateError('Cannot add items to a paid order');
    _items.add(item);
  }

  int totalPrice() => _items.fold(0, (sum, item) => sum + item.price);
  bool isEmpty() => _items.isEmpty;
  void markAsPaid() => _isPaid = true;
}

class Customer {
  int _balance;

  int get balance => _balance;

  Customer(this._balance) {
    if (_balance < 0) throw ArgumentError('Balance cannot be negative');
  }

  bool canPay(Order order) => _balance >= order.totalPrice();

  void pay(Order order) {
    if (order.isEmpty()) throw ArgumentError('Cannot pay an empty order');
    if (!canPay(order)) throw StateError('Insufficient balance');
    _balance -= order.totalPrice();
    order.markAsPaid();
  }

  void deposit(int amount) {
    if (amount <= 0) throw ArgumentError('Amount must be positive');
    _balance += amount;
  }
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
  try {
    customer2.pay(order2);
  } catch (e) {
    print('Error: $e');
  }
  print('Paid: ${order2.isPaid} | Balance: ${customer2.balance}');

  print('=== Scenario 3: Empty order ===');
  final order3 = Order();
  final customer3 = Customer(100);
  try {
    customer3.pay(order3);
  } catch (e) {
    print('Error: $e');
  }

  print('=== Scenario 4: Invalid item ===');
  try {
    final bad = MenuItem('', -5);
    print(bad.describe());
  } catch (e) {
    print('Error: $e');
  }

  print('=== Edge 5: Add to paid order ===');
  final order5 = Order();
  order5.addItem(MenuItem('Latte', 45));
  Customer(100).pay(order5);
  try {
    order5.addItem(MenuItem('Espresso', 35));
  } catch (e) {
    print('Error: $e');
  }

  print('=== Edge 6: Deposit then check ===');
  final c = Customer(50);
  c.deposit(50);
  print('Balance after deposit: ${c.balance}');

  print('=== Edge 7: Invalid deposit (0) ===');
  try {
    c.deposit(0);
  } catch (e) {
    print('Error: $e');
  }

  print('=== Edge 8: Discount applied ===');
  final latte = MenuItem('Latte', 45);
  latte.applyDiscount(10);
  print('Latte after 10% off: ${latte.price}');
}