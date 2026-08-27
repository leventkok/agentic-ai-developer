import { describe, it, beforeEach } from "node:test";
import assert from "node:assert/strict";
import { retryWithBackoff } from "../src/async/retry.js";
import { debounceAsync } from "../src/async/debounce.js";
import { fetchUserWorkflow, resetFlakyAttempts, setFlakyFailures } from "../src/services/fetch-workflow.js";
import { drainPending } from "../src/async/shutdown.js";

describe("retryWithBackoff", () => {
  it("retries transient failures", async () => {
    let calls = 0;
    const value = await retryWithBackoff(
      async () => {
        calls++;
        if (calls < 3) throw new Error("transient");
        return "ok";
      },
      { maxAttempts: 4, baseDelayMs: 1 }
    );
    assert.equal(value, "ok");
    assert.equal(calls, 3);
  });
});

describe("debounceAsync", () => {
  it("only resolves last call", async () => {
    let calls = 0;
    const debounced = debounceAsync(async (n: number) => {
      calls++;
      return n;
    }, 20);

    debounced(1);
    debounced(2);
    assert.equal(await debounced(3), 3);
    await new Promise((r) => setTimeout(r, 30));
    assert.equal(calls, 1);
  });
});

describe("fetchUserWorkflow", () => {
  beforeEach(() => resetFlakyAttempts());

  it("composes retry + Result for flaky fetch", async () => {
    setFlakyFailures(2);
    const result = await fetchUserWorkflow("1");
    assert.equal(result.ok, true);
    if (result.ok) assert.equal(result.value.name, "Ada");
  });

  it("returns not_found without retrying", async () => {
    const result = await fetchUserWorkflow("999");
    assert.equal(result.ok, false);
    if (!result.ok) assert.equal(result.error.kind, "not_found");
  });
});

describe("shutdown", () => {
  it("drainPending clears tracked promises", async () => {
    const slow = new Promise<string>((resolve) => setTimeout(() => resolve("done"), 10));
    const { trackPromise } = await import("../src/async/shutdown.js");
    void trackPromise(slow);
    await drainPending();
    assert.equal(await slow, "done");
  });
});
