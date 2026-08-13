/*
 DAY 11 — Inheritance
 Hierarchy: Notification → EmailNotification, SmsNotification (1 level only)
*/


class Notification {
  final String message;
  final String recipient;

  Notification(this.message, this.recipient){
    if(message.isEmpty) throw ArgumentError('Message cannot be empty');
    if(recipient.isEmpty) throw ArgumentError('Recipient cannot be empty');
  }

  void send() {
    print('Sending notification to $recipient: $message');
  }


  String describe() => 'Notification(message: $message, recipient: $recipient)';

}


class EmailNotification extends Notification {
  final String email;

  EmailNotification(this.email, String message) : super(message, email);

  @override
  void send() {
    print('Sending email to $email: Message: $message');
  }
}


class SmsNotification extends Notification {
  final String phone;

  SmsNotification(this.phone, String message) : super(message, phone);

  @override
  void send() {
    print('Sending SMS to $phone: Message: $message');
  }
}


void main(){
  print('=== Email notification ===');
  final email = EmailNotification('levent@example.com', 'Your coffee is ready!');
  email.send();
  print(email.describe());


  print('=== SMS notification ===');
  final sms = SmsNotification('+1234567890', 'Your coffee is ready!');
  sms.send();
  print(sms.describe());


  print('=== Base Notification ===');
  final base = Notification('staff', 'New order received!');
  base.send();

}