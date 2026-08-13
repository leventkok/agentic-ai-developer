/*
 DAY 12 — Method Overriding & Polymorphism

 CreditCardPayment.pay — adds card charge step after base validation
 WalletPayment.pay     — adds wallet deduction step after base validation

 Polymorphism demo: List<PaymentMethod> holds mixed types;
   calling pay() runs the correct override at runtime.
*/

class PaymentMethod{
  final String owner;

  PaymentMethod(this.owner){
    if(owner.isEmpty) throw ArgumentError('Owner cannot be empty');
  }

  void pay(int amount){
    if(amount <= 0) throw ArgumentError('Amount must be greater than 0');
    print('Paying $amount to $owner');
  }

  String describe() => 'PaymentMethod(owner: $owner)';



}


class CreditCardPayment extends PaymentMethod{
  final String lastFourDigits;

  CreditCardPayment(this.lastFourDigits, String owner) : super(owner);

  @override
  void pay(int amount){
    super.pay(amount);
    print('Paying with credit card ending in $lastFourDigits');
  }
}


class WalletPayment extends PaymentMethod{
  WalletPayment(String owner) : super(owner);

  @override
  void pay(int amount){
    super.pay(amount);
    print('Paying with wallet');
  }
}



void main(){
  final methods =<PaymentMethod>[
    CreditCardPayment('1234', 'John Doe'),
    WalletPayment('Jane Smith'),
    PaymentMethod('Alice Johnson'),
  ];

  print('=== Polymorphism: send all methods ===');


  for (final method in methods){
    method.pay(50);
    print(method.describe());
    print('---');
  }


}


