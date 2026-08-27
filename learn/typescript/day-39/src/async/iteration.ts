import type { User } from "../types/user.js";
import { fetchUser } from "../promises/basic.js";

export async function fetchUsersInOrder(ids: string[]): Promise<User[]> {
  const users: User[] = []; 

  for (const id of ids) {
    const user: User = await fetchUser(id);
    users.push(user);
  }

  return users;
}