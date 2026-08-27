import type { User } from "../types/user.js";

const fakeDb: Record<string, User> = {
  "1": { id: "1", name: "Ada", email: "ada@dev" },
};

function delay(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export async function fetchUser(id: string): Promise<User> {
  await delay(10);
  const user = fakeDb[id];
  if (!user) throw new Error("User not found");
  return user;
}

export async function logUser(id: string): Promise<void> {
  const user = await fetchUser(id);
  console.log(`User: ${user.name}`);
}

export function fetchUserName(id: string): Promise<string> {
  return fetchUser(id)
    .then((user) => user.name)
    .catch((err: unknown) => {
      const msg = err instanceof Error ? err.message : "Unknown error";
      throw new Error(`fetchUserName failed: ${msg}`);
    });
}