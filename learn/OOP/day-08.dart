/*
 DAY 8 — Abstraction

 Public API (what callers use):
   - CoffeeShop.processCheckout()
   - Customer.checkout()
   - Order.placeOrder()

 Hidden details (callers don't need):
   - balance subtraction math
   - _items list management
   - fold for totalPrice
*/



abstract class Payable{
  int amountDue() => 0;
  bool get isPaid;
  void markAsPaid();

}

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

class Order implements Payable{
  final List<MenuItem> _items = [];
  bool _isPaid = false;

  
  List<MenuItem> get items => List.unmodifiable(_items);


  Order();

  void addItem(MenuItem item) {
    if (_isPaid) throw StateError('Cannot add items to a paid order');
    _items.add(item);
  }

  int totalPrice() => _items.fold(0, (sum, item) => sum + item.price);
  bool isEmpty() => _items.isEmpty;

  int checkoutTotal() => totalPrice();
  


  void placeOrder(MenuItem item) {
    addItem(item);
  }

  @override
  int amountDue() => totalPrice();

  @override
  bool get isPaid => _isPaid;

  @override
  void markAsPaid() => _isPaid = true;


}

class Customer {
  int _balance;

  int get balance => _balance;

  Customer(this._balance) {
    if (_balance < 0) throw ArgumentError('Balance cannot be negative');
  }

  bool canPay(Payable payable) => _balance >= payable.amountDue();

  

  void pay(Payable payable) {
  if (payable.isPaid) throw StateError('Already paid');
  if (payable.amountDue() == 0) throw ArgumentError('Nothing to pay');
  if (_balance < payable.amountDue()) throw StateError('Insufficient balance');
  _balance -= payable.amountDue();
  payable.markAsPaid();
}

  void deposit(int amount) {
    if (amount <= 0) throw ArgumentError('Amount must be positive');
    _balance += amount;
  }

  void checkout(Payable payable) {
    pay(payable);
  }
}


class CoffeeShop {
  void processCheckout(Customer customer, Order order) {
    customer.checkout(order);
    print('Receipt: ${order.checkoutTotal()} TL — Thank you!');
  }
}

void main() {
  final shop = CoffeeShop();

  print('=== Scenario 1: Happy checkout ===');
  final order1 = Order();
  order1.placeOrder(MenuItem('Latte', 45));
  order1.placeOrder(MenuItem('Espresso', 35));
  shop.processCheckout(Customer(100), order1);

  print('=== Scenario 2: Insufficient funds ===');
  final order2 = Order();
  order2.placeOrder(MenuItem('Latte', 45));
  try {
    shop.processCheckout(Customer(30), order2);
  } catch (e) {
    print('Error: $e');
  }

  print('=== Scenario 3: Empty order ===');
  try {
    shop.processCheckout(Customer(100), Order());
  } catch (e) {
    print('Error: $e');
  }
}