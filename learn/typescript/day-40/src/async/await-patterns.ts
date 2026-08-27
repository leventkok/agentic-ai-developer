import { fetchUser } from "../promises/basic.js";

export async function fetchUserNameAsync(id: string): Promise<string> {
  try {
    const user = await fetchUser(id);
    return user.name;
  } catch (err: unknown) {
    const msg = err instanceof Error ? err.message : "Unknown error";
    throw new Error(`fetchUserName failed: ${msg}`);
  }
}