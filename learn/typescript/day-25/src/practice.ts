

type DbUser = {
    id: number;
    name: string;
    email: string;
    passwordHash: string;
    createdAt: Date;
    updatedAt: Date;
};


type PublicUser = Omit<DbUser, "passwordHash">;

type CreateUserInput = Omit<DbUser, "id" | "createdAt" | "updatedAt" | "passwordHash"> & {
    password: string;
};


type UpdateUserInput = Partial<Omit<DbUser, "id" | "createdAt" | "updatedAt">>;



function toPublic(user: DbUser): PublicUser {
    const { passwordHash, ...publicFields } = user;
    return publicFields;
}


function createUser(input: CreateUserInput): DbUser {
    return {
        id: 1,
        name: input.name,
        email: input.email,
        passwordHash: `hash:${input.password}`,
        createdAt: new Date(),
        updatedAt: new Date(),
    };
}


type Environment = "dev" | "staging" | "prod";

type Config = {
    apiUrl: string;
    debug: boolean;
}

type ConfigMap = Record<Environment, Config>;

const configs: ConfigMap = {
    dev: {apiUrl: "http://localhost:3000", debug: true},
    staging: {apiUrl: "https://staging.api.cafe.com", debug: true},
    prod: {apiUrl: "https://api.cafe.com", debug: false},
};

function getConfig(env: Environment): Config {
    return configs[env];
}

function withLogging<F extends (...args: any[]) => Promise<any>>(
  fn: F,
  label: string
): (...args: Parameters<F>) => ReturnType<F> {
  return (async (...args: Parameters<F>) => {
    console.log(`[${label}] called with`, args);
    const result = await fn(...args);
    console.log(`[${label}] returned`, result);
    return result;
  }) as (...args: Parameters<F>) => ReturnType<F>;
}

async function fetchMenu(id: number): Promise<{ id: number; name: string }> {
  return { id, name: "Latte" };
}

const loggedFetchMenu = withLogging(fetchMenu, "fetchMenu");

// Type audit (Day 22): MenuItemDTO was Pick<MenuItem, ...> — if we duplicated
// { name, price, category } manually, it would drift when MenuItem changes.
// Same pattern here: DbUser is the single source of truth for all user types.

// Run demos
const dbUser = createUser({
  name: "Ali",
  email: "ali@cafe.com",
  password: "secret",
});
console.log("Public:", toPublic(dbUser).name);
console.log("Prod config:", getConfig("prod").apiUrl);

const patch: UpdateUserInput = { name: "Ali K." };
console.log("Patch:", patch);

loggedFetchMenu(1).then((menu) => console.log("Menu:", menu.name));