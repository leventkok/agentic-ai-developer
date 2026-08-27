import type { User } from "../types/user.js";
import type { ApiError } from "../types/errors.js";
import type { Result } from "../types/result.js";
import { ok, err } from "../types/result.js";
import { fetchUserCancellable } from "./cancellable-user.js";
import { retryWithBackoff } from "../async/retry.js";
import { isAbortError } from "../utils/abort-error.js";
import { trackPromise } from "../async/shutdown.js";

let flakyAttempts = 0;

/** Test hook — reset flaky counter between tests */
export function resetFlakyAttempts(): void {
  flakyAttempts = 0;
}

/** Simulates transient network failure on first calls */
export function setFlakyFailures(count: number): void {
  flakyAttempts = -count;
}

async function fetchWithTransientRetries(
  id: string,
  signal?: AbortSignal
): Promise<Result<User, ApiError>> {
  return retryWithBackoff(
    async () => {
      if (signal?.aborted) {
        throw new DOMException("Aborted", "AbortError");
      }

      flakyAttempts++;
      if (flakyAttempts <= 0) {
        throw new Error("transient network failure");
      }

      const result = await fetchUserCancellable(id, signal);
      if (!result.ok) {
        if (result.error.kind === "network") {
          throw new Error("transient network failure");
        }
        return result;
      }
      return result;
    },
    {
      maxAttempts: 4,
      baseDelayMs: 5,
      isRetryable: (e) =>
        e instanceof Error && e.message.includes("transient"),
    }
  );
}

/** Composed workflow: track → retry → abort-aware Result fetch */
export async function fetchUserWorkflow(
  id: string,
  signal?: AbortSignal
): Promise<Result<User, ApiError>> {
  try {
    return await trackPromise(fetchWithTransientRetries(id, signal));
  } catch (e: unknown) {
    if (isAbortError(e)) {
      return err({ kind: "network", message: "Request cancelled" });
    }
    throw e;
  }
}
