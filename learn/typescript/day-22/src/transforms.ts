

type MenuItem = {
    id: number;
    name: string;
    price: number;
    cost: number;
    category: string;
};

type MenuItemDTO  =  Pick<MenuItem, "name" | "price" | "category">;


function toDTO(item: MenuItem): MenuItemDTO {
    return {name: item.name, price: item.price, category: item.category};
}

type User ={
    id: number;
    name: string;
    email: string;
    passwordHash: string;
}


type PublicUser = Omit<User, "passwordHash">;

function toPublicUser(user: User): PublicUser {
    const {passwordHash, ...publicFields} = user;
    return publicFields;
}

type Role = "admin" | "barista" | "customer";


type Permission = "manage_menu" | "view_orders" | "place_order";


type RolePermissions = Record<Role, Permission[]>;

const permissions: RolePermissions = {
    admin: ["manage_menu", "view_orders", "place_order"],
    barista: ["view_orders"],
    customer: ["place_order"],
}


// type MenuSummary = { name: string; price: number };  // delete this

type MenuSummary = Pick<MenuItem, "name" | "price">;
// Pick/Omit derive from MenuItem — if domain changes, DTOs update automatically


const latte: MenuItem = { id: 1, name: "Latte", price: 45, cost: 15, category: "hot" };
console.log("DTO:", toDTO(latte));

const user: User = { id: 1, name: "Ali", email: "ali@cafe.com", passwordHash: "secret" };
console.log("Public:", toPublicUser(user));

console.log("Admin perms:", permissions.admin);
console.log("Summary type demo:", { name: "Espresso", price: 30 } satisfies MenuSummary);