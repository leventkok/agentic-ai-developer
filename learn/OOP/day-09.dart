/*
 DAY 9 — Mutable vs Immutable
 Mutable (MenuItem):
   + Kolay kullanım, tek nesne
   - Paylaşınca yan etki riski (başkası değiştirebilir)
   - Debug zor: "kim fiyatı değiştirdi?"
 Immutable (ImmutableMenuItem):
   + Güvenli paylaşım, öngörülebilir
   + Flutter/state management için ideal
   - Her değişimde yeni nesne = daha fazla allocation
 RotaCoffee seçimi:
   - MenuItem (fiyat) → immutable value object
   - Customer, Order → mutable entity (identity var, state değişir)
*/

abstract class Payable {
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
      throw ArgumentError('Discount must be between 0 and 50');
    }
    _price = _price - (_price * percent ~/ 100); // same object, price changes
  }

  String describe() => 'MenuItem(name: $name, price: $price)';
}

class ImmutableMenuItem {
  final String name;
  final int price;

  const ImmutableMenuItem(this.name, this.price)
      : assert(name != ''),
        assert(price >= 0);

  ImmutableMenuItem withDiscount(int percent) {
    if (percent < 0 || percent > 50) {
      throw ArgumentError('Discount must be between 0 and 50');
    }
    final newPrice = price - (price * percent ~/ 100);
    return ImmutableMenuItem(name, newPrice); // new object
  }

  String describe() => 'ImmutableMenuItem(name: $name, price: $price)';
}

class Order implements Payable {
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

  void placeOrder(MenuItem item) => addItem(item);

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

  void checkout(Payable payable) => pay(payable);
}

class CoffeeShop {
  void processCheckout(Customer customer, Order order) {
    customer.checkout(order);
    print('Receipt: ${order.checkoutTotal()} TL — Thank you!');
  }
}

void main() {

  print('=== Mutable: applyDiscount ===');
  final mutable = MenuItem('Latte', 45);
  mutable.applyDiscount(10);
  print('After discount: ${mutable.price}'); 

  print('=== Immutable: withDiscount ===');
  final original = ImmutableMenuItem('Latte', 45);
  final discounted = original.withDiscount(10);
  print('Original: ${original.price}');   
  print('Discounted: ${discounted.price}'); 
  print('Same object? ${identical(original, discounted)}'); 

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