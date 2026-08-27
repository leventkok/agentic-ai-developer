import { describe, it } from "node:test";
import assert from "node:assert/strict";
import { ok, err } from "../src/types/result.js";
import { fetchUserWithResult } from "../src/services/user-service.js";
import { errorMessage } from "../src/utils/safe-catch.js";

describe("Result", () => {
  it("ok and err shapes", () => {
    assert.equal(ok(1).ok, true);
    assert.equal(err("fail").ok, false);
  });
});

describe("fetchUserWithResult", () => {
  it("returns user on success", async () => {
    const r = await fetchUserWithResult("1");
    assert.equal(r.ok && r.value.name, "Ada");
  });

  it("returns not_found error", async () => {
    const r = await fetchUserWithResult("999");
    assert.equal(r.ok, false);
    if (!r.ok) assert.equal(r.error.kind, "not_found");
  });
});

describe("errorMessage", () => {
  it("handles Error and string", () => {
    assert.equal(errorMessage(new Error("x")), "x");
    assert.equal(errorMessage("y"), "y");
  });
});