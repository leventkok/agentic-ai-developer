// Day 12: Implementing and Extending Contracts

interface Serializable {
    toJSON(): string;
}

class MenuItem implements Serializable {
    constructor(public name: string, public price: number) {}

    toJSON(): string {
        return JSON.stringify({ name: this.name, price: this.price });
    }
}


interface Identifiable {
    id: number;
}

interface Printable {
    print(): void;
}


class Order implements Identifiable, Printable {
    constructor(public id: number, public item: string, public total: number) {}

    print(): void {
        console.log(`Order ${this.id}: ${this.item} - ${this.total} TL`);
    }
}


interface Entity {
    id: number;
}

interface Timestamped extends Entity {
    createdAt: Date;
}

interface Auditable extends Timestamped {
    createdBy: string;
}

const auditLog: Auditable = {
    id: 1,
    createdAt: new Date(),
    createdBy: "John Doe",
};

// Task 4: missing createdBy would fail compile — compiler enforces the contract
const fixedAudit: Auditable = {
    id: 2,
    createdAt: new Date(),
    createdBy: "Jane",
};

// Run demos
const latte = new MenuItem("Latte", 4.5);
console.log(latte.toJSON());

const order = new Order(101, "Latte", 45);
order.print();

console.log("Audit:", auditLog.createdBy);
console.log("Fixed audit:", fixedAudit.createdBy);