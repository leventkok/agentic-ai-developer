// TASK 1

class User {
    final String name;
    final String email;
    final int age;

    User(this.name, this.email, this.age);

    // TASK 3
    String describeUser(){
        return "User(name: $name, email: $email, age: $age)";
    }


}

void main(){
    // TASK 2
    final item1 = User(name: "John", email: "john@example.com", age: 20);
    final item2 = User(name: "Jane", email: "jane@example.com", age: 21);
    final item3 = User(name: "Jim", email: "jim@example.com", age: 22);

    // TASK 4
    print('Item 1: ${item1.name} - ${item1.email} - ${item1.age}');
    print('Item 2: ${item2.name} - ${item2.email} - ${item2.age}');

    print(item3.describeUser());
}