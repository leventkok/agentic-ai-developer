import { describe, it } from "node:test";
import assert from "node:assert/strict";
import { isAbortError } from "../src/utils/abort-error.js";
import { fetchUserCancellable } from "../src/services/cancellable-user.js";
import { withTimeout } from "../src/async/timeout.js";

describe("isAbortError", () => {
  it("detects AbortError", () => {
    assert.equal(isAbortError(new DOMException("x", "AbortError")), true);
    assert.equal(isAbortError(new Error("fail")), false);
  });
});

describe("fetchUserCancellable", () => {
  it("cancels when signal aborted", async () => {
    const controller = new AbortController();
    const pending = fetchUserCancellable("1", controller.signal);
    controller.abort();
    const result = await pending;
    assert.equal(result.ok, false);
    if (!result.ok) {
      assert.equal(result.error.kind, "network");
    }
  });

  it("returns user when not cancelled", async () => {
    const result = await fetchUserCancellable("1");
    assert.equal(result.ok, true);
    if (result.ok) {
      assert.equal(result.value.name, "Ada");
    }
  });
});

describe("withTimeout", () => {
  it("aborts slow work", async () => {
    await assert.rejects(
      () =>
        withTimeout(async (signal) => {
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
        }, 20),
      (err: unknown) => isAbortError(err)
    );
  });
});
