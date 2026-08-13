/*
 DAY 14 — Composition over Inheritance

 ❌ BAD (forced inheritance):
   EmailOrder extends Order, PushOrder extends Order
   Why awkward:
   - An order is not a type of email or push notification
   - Adding SMS means another subclass
   - Order logic duplicated across subclasses

 ✅ GOOD (composition):
   Order HAS-A ReceiptSender — swap sender without subclassing
*/


abstract class ReceiptSender {
  void send(String orderId, String summary);

}

class EmailSender extends ReceiptSender {
  final String email;
  EmailSender(this.email);
  @override
  void send(String orderId, String summary) {
    print('Sending email to $email for order $orderId: $summary');
  }
}


class PushSender extends ReceiptSender {
  final String deviceId;
  PushSender(this.deviceId);
  @override
  void send(String orderId, String summary) {
    print('Sending push notification to $deviceId for order $orderId: $summary');
  }
}


class Order{
  final String id;
  final String summary;
  final ReceiptSender _notifier;

  Order(this.id, this.summary, this._notifier);

  void complete(){
    print('Order $id completed');
    _notifier.send(id, summary);
  }
}


/*
 Rule of thumb:
   Inheritance → true is-a (Notification → EmailNotification)
   Composition → has-a / swap behavior (Order has ReceiptSender)
*/


void main() {
  print('=== Order With Email Sender ===');
  Order('RC-1', 'Latte 45 TL', EmailSender('levent@mail.com')).complete();

  print('=== Same Order logic, Push sender ===');
  Order('RC-2', 'Espresso 35 TL', PushSender('device-xyz')).complete();
}