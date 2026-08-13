
/*
 DAY 17 — LSP, ISP, DIP

 LSP fixed: CoffeeBrewer only requires brew() — no throwing overrides
 ISP fixed: Brewable + Cleanable instead of fat CoffeeStation
 DIP fixed: CheckoutService depends on PaymentGateway, not StripeGateway

 Violations we avoided:
   LSP — InstantMachine throwing on grindBeans
   ISP — SimpleDrip forced to implement diagnostics
   DIP — CheckoutService hard-coded to Stripe
*/







// ❌ LSP violation:
// class InstantMachine extends CoffeeBrewer {
//   @override
//   void grindBeans() => throw UnsupportedError('No grinder!');
// }
// Caller expects ALL CoffeeBrewers to grind — InstantMachine breaks that.


abstract class CoffeeBrewer{
  String get name;
  void brew();

}

class EspressoMachine extends CoffeeBrewer{
  @override
  // TODO: implement name
  String get name => 'Espresso';

  @override
  void brew() => print('$name: grinding + extracting');
}


class FilterMachine extends CoffeeBrewer{
  @override
  String get name => 'Filter';

  @override
  void brew() => print('$name: drip brewing');


  
}


void serveCoffe(CoffeeBrewer machine){
    print('Using ${machine.name}');
    machine.brew();    
}


// ❌ Fat interface — forces SimpleDrip to implement diagnostics + receipts
// abstract class CoffeeStation {
//   void brew();
//   void clean();
//   void runDiagnostics();
//   void printReceipt();
// }


abstract class Brewable{
  void brew();

}

abstract class Cleanable{
  void clean();
}


class SimpleDripMachine implements Brewable, Cleanable{
  @override
  void brew() => print('Drip Brewing');


  @override
  void clean() => print('Rinsing filter');

}


class ProEspressoMachine implements Brewable, Cleanable{

  @override
  void brew() => print('Espresso extraction');

  @override 
  void clean() => print('Backflush group head');
}

void dailyClean(Cleanable machine) => machine.clean();

// ❌ CheckoutService creates StripeGateway directly — hard to test/swap
// class CheckoutService {
//   void pay(int amount) => StripeGateway().charge(amount);
// }


abstract class PaymentGateway{
  bool charge(int amount);

}


class StripeGateway implements PaymentGateway{
  @override
  bool charge(int amount) {
    print('Stripe: charging \$ $amount TL');
    return true;
  }
}

class MockGateway implements PaymentGateway{
  @override
  bool charge(int amount){
    print('Mock: charged $amount TL (Test mode)');
    return true;
  }
}



class CheckoutService{
  final PaymentGateway _gateway;
  CheckoutService(this._gateway);

  void checkout(int amount){
    if(!_gateway.charge(amount)){
      throw StateError('Payment failed');
    }
    print('Order completed successfully');
  }
}


void main() {
  print('=== LSP: substitutable brewers ===');
  for (final machine in [EspressoMachine(), FilterMachine()]) {
    serveCoffe(machine);
  }

  print('\n=== ISP: clean only what supports it ===');
  dailyClean(SimpleDripMachine());

  print('\n=== DIP: swap gateway without changing checkout ===');
  CheckoutService(StripeGateway()).checkout(80);
  CheckoutService(MockGateway()).checkout(80);
}