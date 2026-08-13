

class MenuItem {
  final String name;
  final int price;
  const MenuItem(this.name, this.price);
}

class Order {
  final String id;
  final List<MenuItem> _items = [];

  Order(this.id);

  List<MenuItem> get items => List.unmodifiable(_items);

  void addItem(MenuItem item) {
    if (item.price < 0) throw ArgumentError('Price cannot be negative');
    _items.add(item);
  }

  bool isEmpty() => _items.isEmpty;
}

class DrinkFactory {
  static MenuItem create(String type) {
    switch (type) {
      case 'latte':
        return const MenuItem('Latte', 45);
      case 'espresso':
        return const MenuItem('Espresso', 35);
      case 'filter':
        return const MenuItem('Filter', 30);
      default:
        throw ArgumentError('Unknown drink: $type');
    }
  }
}

class OrderCalculator {
  int calculateTotal(Order order) =>
      order.items.fold(0, (sum, item) => sum + item.price);
}

abstract class DiscountStrategy {
  int apply(int price);
  String get label;
}

class StudentDiscount implements DiscountStrategy {
  @override
  int apply(int price) => price - (price * 15 ~/ 100);
  @override
  String get label => 'Student 15% off';
}

class LoyaltyDiscount implements DiscountStrategy {
  @override
  int apply(int price) => price - (price * 10 ~/ 100);
  @override
  String get label => 'Loyalty 10% off';
}

abstract class PaymentGateway {
  bool charge(String orderId, int amount);
}

class CashPayment implements PaymentGateway {
  @override
  bool charge(String orderId, int amount) {
    print('Cash: charged $amount TL for $orderId');
    return true;
  }
}

class FailingPayment implements PaymentGateway {
  @override
  bool charge(String orderId, int amount) {
    print('Payment declined for $orderId');
    return false;
  }
}

class ReceiptPrinter {
  void printReceipt(Order order, int total, String discountLabel) {
    print('--- Receipt: ${order.id} ---');
    for (final item in order.items) {
      print('  ${item.name}: ${item.price} TL');
    }
    print('  Discount: $discountLabel');
    print('  Total: $total TL');
  }
}

abstract class OrderObserver {
  void onStatusChanged(String orderId, String status);
}

class OrderStatus {
  final String orderId;
  final _observers = <OrderObserver>[];
  OrderStatus(this.orderId);

  void subscribe(OrderObserver o) => _observers.add(o);

  void updateStatus(String status) {
    for (final o in _observers) {
      o.onStatusChanged(orderId, status);
    }
  }
}

class BaristaDisplay implements OrderObserver {
  @override
  void onStatusChanged(String orderId, String status) =>
      print('Barista: order #$orderId is $status');
}

class CustomerNotifier implements OrderObserver {
  @override
  void onStatusChanged(String orderId, String status) =>
      print('Customer: your order #$orderId is $status');
}

class RotaCoffeeShop {
  final PaymentGateway _payment;
  final OrderCalculator _calculator;
  final ReceiptPrinter _printer;

  RotaCoffeeShop({
    required PaymentGateway payment,
    required OrderCalculator calculator,
    required ReceiptPrinter printer,
  })  : _payment = payment,
        _calculator = calculator,
        _printer = printer;

  void checkout(Order order, DiscountStrategy discount) {
    if (order.isEmpty()) {
      print('Cannot checkout an empty order');
      return;
    }

    final total = discount.apply(_calculator.calculateTotal(order));

    if (!_payment.charge(order.id, total)) {
      print('Payment failed — no receipt, order not ready');
      return;
    }

    _printer.printReceipt(order, total, discount.label);

    final status = OrderStatus(order.id)
      ..subscribe(BaristaDisplay())
      ..subscribe(CustomerNotifier());
    status.updateStatus('PAID');
    status.updateStatus('READY');
  }
}

void main() {
  final shop = RotaCoffeeShop(
    payment: CashPayment(),
    calculator: OrderCalculator(),
    printer: ReceiptPrinter(),
  );

  print('=== Order 1: student discount ===');
  final order1 = Order('RC-001');
  order1.addItem(DrinkFactory.create('latte'));
  order1.addItem(DrinkFactory.create('espresso'));
  shop.checkout(order1, StudentDiscount());

  print('\n=== Order 2: loyalty discount, payment fails ===');
  final failShop = RotaCoffeeShop(
    payment: FailingPayment(),
    calculator: OrderCalculator(),
    printer: ReceiptPrinter(),
  );
  final order2 = Order('RC-002');
  order2.addItem(DrinkFactory.create('filter'));
  failShop.checkout(order2, LoyaltyDiscount());
}




/*
 REFLECTION — 20 days of OOP

 Learned:
   1. Encapsulation — keeping data private and changing it only through
      methods with rules (invariants) so objects can't end up in a broken state.
   2. The SOLID principles finally clicked, especially SRP (one class, one job)
      and DIP (depend on abstractions like PaymentGateway, not concrete classes).
   3. Design patterns like Factory, Strategy, and Observer, and when each one
      actually helps instead of forcing them everywhere.

 Will practice next:
   1. Applying OOP and SOLID in other languages, not just Dart.
   2. Using these ideas in real projects, since I think the whole structure
      will be useful going forward.
*/