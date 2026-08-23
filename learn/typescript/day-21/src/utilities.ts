

type User = {
    id: number;
    name: string;
    email: string;
};

function updateUser(user: User, patch: Partial<User>): User{
    return {...user, ...patch};
}

const ali: User ={id: 1, name: "Ali", email: "ali@example.com"};
const updated = updateUser(ali, {name: "Ali K."});


type ApiUser = {
    id?: number;
    name?: string;
    email?: string;
};


function toDomainUser (raw: ApiUser): Required<ApiUser> {
    if(!raw.id || !raw.name || !raw.email) {
        throw new Error("Missing required fields");
    }
    return {id: raw.id, name: raw.name, email: raw.email};
}



type ShopConfig = Readonly<{
    shopName: string;
    apiUrl: string;
  }>;
  const config: ShopConfig = {
    shopName: "MasterFabric Cafe",
    apiUrl: "https://api.masterfabric.cafe",
};

type LockedPartialUser = Partial<Readonly<User>>;

const patch: LockedPartialUser = { name: "Zeynep" };



console.log("Updated:", updated.name);
console.log("Domain:", toDomainUser({ id: 2, name: "Zeynep", email: "z@cafe.com" }));
console.log("Config:", config.shopName);
console.log("Patch:", patch);