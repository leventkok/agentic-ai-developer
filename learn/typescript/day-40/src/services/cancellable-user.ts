import type { User } from "../types/user.js";
import type { ApiError } from "../types/errors.js";
import type { Result } from "../types/result.js";
import { ok, err } from "../types/result.js";
import { isAbortError } from "../utils/abort-error.js";

const fakeDb: Record<string, User> = {
  "1": { id: "1", name: "Ada", email: "ada@dev" },
};

function delay(ms: number, signal?: AbortSignal): Promise<void> {
  return new Promise((resolve, reject) => {
    if (signal?.aborted) {
      reject(new DOMException("Aborted", "AbortError"));
      return;
    }

    const timer = setTimeout(resolve, ms);
    signal?.addEventListener(
      "abort",
      () => {
        clearTimeout(timer);
        reject(new DOMException("Aborted", "AbortError"));
      },
      { once: true }
    );
  });
}

async function loadFromDb(id: string, signal?: AbortSignal): Promise<User | null> {
  await delay(50, signal);
  if (signal?.aborted) return null;
  return fakeDb[id] ?? null;
}

export async function fetchUserCancellable(
  id: string,
  signal?: AbortSignal
): Promise<Result<User, ApiError>> {
  try {
    if (!id.trim()) {
      return err({ kind: "validation", field: "id", message: "id is required" });
    }

    const user = await loadFromDb(id, signal);
    if (!user) {
      return err({ kind: "not_found", resource: "user", id });
    }
    return ok(user);
  } catch (e: unknown) {
    if (isAbortError(e)) {
      return err({ kind: "network", message: "Request cancelled" });
    }
    throw e;
  }
}
