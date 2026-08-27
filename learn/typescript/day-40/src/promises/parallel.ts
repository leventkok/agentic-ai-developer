import type { User } from "../types/user.js";
import { fetchUser } from "./basic.js";

function fetchPosts(userId: string): Promise<string[]> {
  return Promise.resolve([`post-a-by-${userId}`, `post-b-by-${userId}`]);
}

export async function fetchUserAndPosts(
  userId: string
): Promise<[User, string[]]> {
  return Promise.all([fetchUser(userId), fetchPosts(userId)]);
}

export function fetchAllSettled(
  ids: string[]
): Promise<PromiseSettledResult<User>[]> {
  return Promise.allSettled(ids.map((id) => fetchUser(id)));
}
