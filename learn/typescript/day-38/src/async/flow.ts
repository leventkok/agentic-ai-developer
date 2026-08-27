import type { User } from "../types/user.js";
import { fetchUser } from "../promises/basic.js";

export async function fetchUsersSequential(ids: string[]): Promise<User[]> {
  const users: User[] = [];
  for (const id of ids) {
    users.push(await fetchUser(id));
  }
  return users;
}

export async function fetchUsersParallel(ids: string[]): Promise<User[]> {
  return Promise.all(ids.map((id) => fetchUser(id)));
}
