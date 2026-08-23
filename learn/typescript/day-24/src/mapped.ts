type User = {
    id: number;
    name: string;
    email: string;
};


type Reconstructed = { [K in keyof User]: User[K] };

const ali: Reconstructed = {id: 1, name: "Ali", email: "ali@cafe.com"};


type Nullable<T> = {
    [K in keyof T]: T[K] | null;
}

type NullableUser = Nullable<User>;

const draft: NullableUser = {id: 1, name: null, email: null};



type Immutable<T> = {
    readonly [K in keyof T]: T[K];
};

type AllOptional<T> = {
    [K in keyof T]?: T[K];
};

type Mutable<T> = {
    -readonly [K in keyof T]: T[K];
};


type AllRequired<T> = {
    [K in keyof T]-?: T[K];
};


const config: Immutable<User> = {id: 1, name: "Ali", email: "ali@cafe.com"};

const patch: AllOptional<User> = {name: "Zeynep"};


type MyPartial<T> = {
    [K in keyof T]?: T[K];
}

// Partial maps over keyof T and adds ? to each property — that's why all fields become optional



console.log("User:", ali.name);
console.log("Draft:", draft);
console.log("Patch:", patch);

const partialUser: MyPartial<User> = { name: "Latte" };
console.log("Partial:", partialUser);