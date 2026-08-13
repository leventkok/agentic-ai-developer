

/*
 DAY 15 — Phase 3 Capstone

 Inheritance:  Reward → FreeDrinkReward, DiscountReward (true is-a)
 Contract:     Reward.applyTo(), ReceiptSender.send()
 Polymorphism: List<Reward> processed in redeemAll loop
 Composition:  LoyaltyProgram HAS-A ReceiptSender (not EmailLoyaltyProgram extends ...)

 Critique — avoided deep inheritance:
   Did NOT make EmailLoyaltyProgram extends LoyaltyProgram.
   Notification channel is swappable via composition instead.
*/

abstract class Reward{
  String get name;
  String applyTo(String orderSummary);

}

class FreeDrinkReward implements Reward{
  @override
  String get name => 'Free Drink';

  @override 
  String applyTo(String orderSummary) {
    return '$orderSummary + Free Latte (0 TL)';

  }
  
}



class DiscountReward implements Reward{
  final int percent;
  DiscountReward(this.percent);


  @override
  String get name => '$percent% Off';

  @override
  String applyTo(String orderSummary) {
    return '$orderSummary - $percent% Loyalty discount applied';
  }
}


abstract class ReceiptSender {
  void send(String memberName, String message);
}

class EmailSender implements ReceiptSender {
  final String email;
  EmailSender(this.email);
  @override
  void send(String memberName, String message) {
    print('Email to $email: $message');
  }
}

class PushSender implements ReceiptSender {
  final String deviceId;
  PushSender(this.deviceId);
  @override
  void send(String memberName, String message) {
    print('Push to $deviceId: $message');
  }
}


class LoyaltyProgram {
  final String memberName;
  final ReceiptSender _notifier;           
  final List<Reward> _rewards = [];

  LoyaltyProgram(this.memberName, this._notifier);

  void addReward(Reward reward) => _rewards.add(reward);

  void redeemAll(String orderSummary) {
    var summary = orderSummary;
    for (final reward in _rewards) {       
      summary = reward.applyTo(summary);
      print('Redeemed: ${reward.name}');
    }
    _notifier.send(memberName, summary);   
  }
}


void main() {
  print('=== Email member ===');
  final emailProgram = LoyaltyProgram('Levent', EmailSender('levent@mail.com'));
  emailProgram.addReward(FreeDrinkReward());
  emailProgram.addReward(DiscountReward(10));
  emailProgram.redeemAll('Order: Latte 45 TL');

  print('\n=== Push member — same program logic ===');
  final pushProgram = LoyaltyProgram('Ayse', PushSender('device-abc'));
  pushProgram.addReward(DiscountReward(20));
  pushProgram.redeemAll('Order: Espresso 35 TL');
}