/*
 DAY 16 — SRP & OCP

 ❌ SMELL — OrderManager (violates SRP):
   - calculateTotal()  → changes when pricing rules change
   - printReceipt()    → changes when receipt format changes
   TWO reasons to change = SRP violation

 ✅ FIX — split into:
   - OrderCalculator  (one job: math)
   - ReceiptPrinter   (one job: output)
*/


class OrderCalculator {
  double calculateTotal(List<int> prices) {
    return prices.fold(0, (sum, p) => sum + p);
  }
}

class ReceiptPrinter {
  void printReceipt(String orderId, double total) {
    print('Receipt for $orderId — Total: $total TL');
  }
}

class OrderService {
  final OrderCalculator _calculator;
  final ReceiptPrinter _printer;

  OrderService(this._calculator, this._printer);

  void checkout(String orderId, List<int> prices) {
    final total = _calculator.calculateTotal(prices);
    _printer.printReceipt(orderId, total);
  }
}

// ❌ OCP violation — adding 'priority' means editing this method:
// double getShippingCost(String type) {
//   switch (type) { case 'standard': return 5; case 'express': return 15; }
// }
// Places to change when adding type: 1 method (+ every caller using strings)


abstract class ShippingStrategy {
  double cost();
  String get label;
}

class StandardShipping implements ShippingStrategy {
  @override
  double cost() => 5;
  @override
  String get label => 'Standard';
}

class ExpressShipping implements ShippingStrategy {
  @override
  double cost() => 15;
  @override
  String get label => 'Express';
}

class PriorityShipping implements ShippingStrategy {
  @override
  double cost() => 25;
  @override
  String get label => 'Priority';
}

class ShippingCalculator {
  final ShippingStrategy _strategy;
  ShippingCalculator(this._strategy);

  double totalWithShipping(double orderTotal) {
    return orderTotal + _strategy.cost();
  }

  String describe() => '${_strategy.label} shipping: ${_strategy.cost()} TL';
}


/*
 BEFORE/AFTER change count (adding Priority shipping):

 Switch approach:  edit getShippingCost switch + update callers → ~2-3 places
 Strategy approach: add PriorityShipping class only           → 1 place

 SRP split: ReceiptPrinter can change format without touching OrderCalculator
*/

void main() {
  print('=== SRP: checkout ===');
  OrderService(OrderCalculator(), ReceiptPrinter())
      .checkout('RC-10', [45, 35]);

  print('\n=== OCP: shipping strategies ===');
  for (final strategy in [StandardShipping(), ExpressShipping(), PriorityShipping()]) {
    final calc = ShippingCalculator(strategy);
    print('${calc.describe()} → total: ${calc.totalWithShipping(80)} TL');
  }
}