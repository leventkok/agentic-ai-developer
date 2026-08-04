class User {
    final String name;
    final String email;
    final int age;
    final String role;

    User(this.name, this.email, this.age, {this.role = 'customer'}) {
        if (name.trim().isEmpty) {
            throw ArgumentError("Name cannot be empty");
        }
        if (age < 0) {
            throw ArgumentError("Age cannot be negative");
        }
    }

    User.guest()
        : name = 'Guest',
          email = 'guest@example.com',
          age = 0,
          role = 'guest';
}

void main() {
    final user1 = User("John", "john@example.com", 20);
    print(user1.name);
    print(user1.age);

    try {
        final user2 = User("", "jane@example.com", -1);
        print(user2.name);
    } catch (e) {
        print("Error: $e");
    }

    final guest = User.guest();
    print(guest.name);
    print(guest.role);
}class User {
    final String name;
    final String email;
    final int age;
    final String role;

    User(this.name, this.email, this.age, {this.role = 'customer'}) {
        if (name.trim().isEmpty) {
            throw ArgumentError("Name cannot be empty");
        }
        if (age < 0) {
            throw ArgumentError("Age cannot be negative");
        }
    }

    User.guest()
        : name = 'Guest',
          email = 'guest@example.com',
          age = 0,
          role = 'guest';
}

void main() {
    final user1 = User("John", "john@example.com", 20);
    print(user1.name);
    print(user1.age);

    try {
        final user2 = User("", "jane@example.com", -1);
        print(user2.name);
    } catch (e) {
        print("Error: $e");
    }

    final guest = User.guest();
    print(guest.name);
    print(guest.role);
}