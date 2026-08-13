


abstract class ReceiptSender{
  void send(String orderId, String sumamry);

  void log(String message) => print('ReceiptSender: $message');

}

class EmailReceiptSender implements ReceiptSender{
  final String email;

  EmailReceiptSender(this.email){
    if(email.isEmpty) throw ArgumentError('Email cannot be empty');
    log('EmailReceiptSender initialized for $email');
  }

  @override
  void send(String orderId, String sumamry){
    print('Sending email receipt to $email for order $orderId: $sumamry');
    log('Email sent successfully');
  }

  @override
  void log(String message) => print('EmailReceiptSender: $message');


}


class PushReceiptSender implements ReceiptSender{
  final String deviceToken;

  PushReceiptSender(this.deviceToken);

  
 

  @override
  void send(String orderId, String sumamry){
    print('Sending push notification to $deviceToken for order $orderId: $sumamry');
    log('Push notification sent successfully');
  }

  @override
  void log(String message) => print('PushRecieptSender: $message');

}




void notifyCustomer(ReceiptSender sender, String orderId, String summary){
  sender.log('Preparing receipt for order $orderId');
  sender.send(orderId, summary);
  sender.log('Customer notified successfully');
}


void main(){
  const orderId = 'RC-42';
  const summary = 'Latte + Espresso = 80 TL';
  print('=== Email receipt ===');
  notifyCustomer(EmailReceiptSender('levent@mail.com'), orderId, summary);
  print('=== Push receipt ===');
  notifyCustomer(PushReceiptSender('device-abc123'), orderId, summary);
}