

/*
Task 1 
Behaviors of before the refactoring:
1. User can review an mobile app,
2. User can see the list of all reviews,
3. User can print a PDF of the reviews (its a summary of the reviews),
4. The app completes of app reviews,

*/

abstract class OrderObserver {
  void onStatusChanged(String orderId, String status);
}

class OrderStatus {
  final String orderId;
  final _observers = <OrderObserver>[];

  OrderStatus(this.orderId);

  void subscribe(OrderObserver o) => _observers.add(o);
  void unsubscribe(OrderObserver o) => _observers.remove(o);

  void updateStatus(String status) {
    for (final o in _observers) {
      o.onStatusChanged(orderId, status);
    }
  }
}

class BaristaDisplay implements OrderObserver {
  @override
  void onStatusChanged(String orderId, String status) {
    print('Barista: order #$orderId is $status');
  }
}

class CustomerNotifier implements OrderObserver {
  @override
  void onStatusChanged(String orderId, String status) {
    print('Customer: your order #$orderId is $status');
  }
}

class KitchenDisplay implements OrderObserver {
  @override
  void onStatusChanged(String orderId, String status) {
    if (status == 'READY') {
      print('Kitchen: order #$orderId is ready to hand out');
    }
  }
}

class Order {
  final String orderId;
  final List<String> items = [];
  final List<int> prices = [];

  Order(this.orderId);

  void addItem(String name, int price) {
    items.add(name);
    prices.add(price);
  }

  int getTotal() => prices.fold(0, (sum, p) => sum + p);
}



class ReceiptPrinter {
  void printReceipt(Order order) {
    print('--- Receipt: ${order.orderId} ---');
    for (var i = 0; i < order.items.length; i++) {
      print('  ${order.items[i]}: ${order.prices[i]} TL');
    }
    print('  Total: ${order.getTotal()} TL');
  }
}


class OrderService {
  final Order order;
  final ReceiptPrinter printer;
  final OrderStatus status;

  OrderService({
    required this.order,
    required this.printer,
    required this.status,
  });

  void checkout() {
    printer.printReceipt(order);
    status.updateStatus('PREPARING');
    status.updateStatus('READY');
  }
}


void main() {
  print('=== Observer demo ===');
  final status = OrderStatus('123');
  final barista = BaristaDisplay();
  status.subscribe(barista);
  status.subscribe(CustomerNotifier());

  status.updateStatus('PREPARING');
  status.updateStatus('READY');

  print('=== After unsubscribing barista ===');
  status.unsubscribe(barista);
  status.updateStatus('PICKED_UP');

  print('\n=== Refactored checkout ===');
  final order = Order('RC-42');
  order.addItem('Latte', 45);
  order.addItem('Espresso', 35);

  final service = OrderService(
    order: order,
    printer: ReceiptPrinter(),
    status: OrderStatus('RC-42')..subscribe(BaristaDisplay())..subscribe(CustomerNotifier()),
  );
  service.checkout();
}