import { fetchUserCancellable } from "./services/cancellable-user.js";
import { withTimeout } from "./async/timeout.js";
import { isAbortError } from "./utils/abort-error.js";
import { formatApiError } from "./utils/safe-catch.js";

export async function runDemo(): Promise<void> {
  const controller = new AbortController();
  const slow = fetchUserCancellable("1", controller.signal);
  setTimeout(() => controller.abort(), 20);

  const cancelled = await slow;
  if (!cancelled.ok) {
    console.log("Cancelled:", formatApiError(cancelled.error));
  }

  try {
    await withTimeout(async (signal) => {
      await new Promise<void>((resolve, reject) => {
        const timer = setTimeout(resolve, 200);
        signal.addEventListener(
          "abort",
          () => {
            clearTimeout(timer);
            reject(new DOMException("Aborted", "AbortError"));
          },
          { once: true }
        );
      });
    }, 30);
  } catch (e: unknown) {
    console.log("Timeout?", isAbortError(e));
  }

  const success = await fetchUserCancellable("1");
  if (success.ok) {
    console.log("Success:", success.value.name);
  }
}

void runDemo();
