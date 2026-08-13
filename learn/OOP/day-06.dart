



class MenuItem{
    final String name;
    int _price;
    int get price => _price;
    MenuItem(this.name, this._price){
        if(name.isEmpty) throw Exception('Name cannot be empty');
        if(price < 0) throw Exception('Price cannot be negative');
    }

    void applyDiscount(int percent){
        if(percent < 0 || percent > 50) throw Exception('Discount percentage must be between 0 and 50');

        _price = _price - (_price * percent ~/ 100);
    }

    String describe() => 'MenuItem(name: $name, price: $price)';
}

class Order{
    final List<MenuItem> _items = [];
    bool _isPaid = false;

    bool get isPaid => _isPaid;

    List<MenuItem> get items => List.unmodifiable(_items);
    Order();


    void addItem(MenuItem item){
        _items.add(item);
    }

    int totalPrice() => _items.fold(0, (sum, item) => sum + item.price);
    bool isEmpty() => _items.isEmpty;
    void markAsPaid(){
        _isPaid = true;
    }


}


class Customer{
    int _balance;

    int get balance => _balance;
    Customer(this._balance){
        if(_balance < 0) throw Exception('Balance cannot be negative');
    }

    void pay(Order order){
        if(order.isEmpty()) return;
        if(_balance < order.totalPrice()) return;
        _balance -= order.totalPrice();
        order.markAsPaid();
    }
    bool canPay(Order order) => _balance >= order.totalPrice();

    void deposit(int amount){
        if(amount < 0) throw Exception('Amount cannot be negative');
        _balance += amount;
    }

}


void main() {

  print('=== Scenario 1: Happy path ===');
  final order1 = Order();
  order1.addItem(MenuItem('Latte', 45));
  order1.addItem(MenuItem('Espresso', 35));
  final customer1 = Customer(100);
  customer1.pay(order1);
  print('Paid: ${order1.isPaid} | Balance left: ${customer1.balance}');


  print('=== Scenario 2: Insufficient funds ===');
  final order2 = Order();
  order2.addItem(MenuItem('Latte', 45));
  final customer2 = Customer(30);        
  customer2.pay(order2);
  print('Paid: ${order2.isPaid} | Balance: ${customer2.balance}');


  print('=== Scenario 3: Empty order ===');
  final order3 = Order();               
  final customer3 = Customer(100);
  customer3.pay(order3);
  print('Paid: ${order3.isPaid} | Balance: ${customer3.balance}');

  print('=== Scenario 4: Invalid item ===');
  try {
    final bad = MenuItem('', -5);        
    print(bad.describe());
  } catch (e) {
    print('Error: $e');
  }



    print('=== Scenario: Deposit and pay ===');
final c = Customer(50);
c.deposit(50);
print('Balance after deposit: ${c.balance}');  // 100

print('=== Scenario: Invalid deposit ===');
try {
  c.deposit(-10);
} catch (e) {
  print('Error: $e');
}

print('=== Scenario: Discount applied ===');
final latte = MenuItem('Latte', 45);
latte.applyDiscount(10);
print('Latte after 10% off: ${latte.price}');  // 40


  //customer.balance = -10;
  // order.isPaid = true;
}
