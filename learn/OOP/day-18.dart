/*
 DAY 18 — Factory & Strategy

 Factory: CoffeeFactory.create() centralizes MenuItem creation
 Strategy: DiscountStrategy swaps discount algorithms without switches

 When NOT to use a pattern:
   A single MenuItem('Water', 0) doesn't need a factory —
   plain constructor is simpler (avoid ceremony).
*/

class MenuItem{
  final String name;
  final int price;

  MenuItem(this.name, this.price);
}

class CoffeFactory{
  static MenuItem create(String type){
    switch(type.toLowerCase()){
      case 'latte':
       return MenuItem('Latte', 45);
      case 'espresso':
       return MenuItem('Espresso', 35);
      case 'americano':
       return MenuItem('Cappuccino', 50);
      default:
       throw ArgumentError('Invalid coffee type: $type');
    }
  }
}

abstract class DiscountStrategy{
  int apply(int price);
  String get label;
}

class StudentDiscount implements DiscountStrategy{
  @override
  int apply(int price) => price - 10;

  @override
  String get label => 'Student (-10 TL)';

  

}


class LoyaltyDiscount implements DiscountStrategy{
  @override
  int apply(int price) => price - (price * 10 ~/100);

  @override
  String get label => 'Loyalty (10%)';
}


class NoDiscount implements DiscountStrategy{
  @override
  int apply(int price) => price;

  @override
  String get label => 'No Discount';
}

void checkout(MenuItem item, DiscountStrategy strategy){
  final finalPrice = strategy.apply(item.price);
  print('${item.name} : ${item.price} TL -> ${finalPrice} TL (${strategy.label})');
}

void main(){
  final latte = CoffeFactory.create('latte');
  final espresso = CoffeFactory.create('espresso');

  print('=== Strategy: different discount ===');
  checkout(latte, StudentDiscount());
  checkout(latte, LoyaltyDiscount());
  checkout(espresso, NoDiscount());

  print('\n=== Factory: unkown drink ===');

  try{
    CoffeFactory.create('mocha');
  }catch(e){
    print('Error: $e');
  }


}
