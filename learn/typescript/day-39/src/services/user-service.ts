import type { User } from "../types/user.js";
import type { ApiError } from "../types/errors.js";
import type { Result } from "../types/result.js";
import { ok, err } from "../types/result.js";

const fakeDb: Record<string, User> = {
  "1": { id: "1", name: "Ada", email: "ada@dev" },
};

function delay(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

/** Result style — expected failures are return values */
export async function fetchUserWithResult(
  id: string
): Promise<Result<User, ApiError>> {
  await delay(10);

  if (!id.trim()) {
    return err({ kind: "validation", field: "id", message: "id is required" });
  }

  const user = fakeDb[id];
  if (!user) {
    return err({ kind: "not_found", resource: "user", id });
  }

  return ok(user);
}

/** Throw style — failures jump to catch */
export async function fetchUserOrThrow(id: string): Promise<User> {
  await delay(10);

  if (!id.trim()) {
    throw new Error("id is required");
  }

  const user = fakeDb[id];
  if (!user) {
    throw new Error(`User ${id} not found`);
  }

  return user;
}
