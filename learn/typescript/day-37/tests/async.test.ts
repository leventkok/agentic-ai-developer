import { describe, it } from "node:test";
import assert from "node:assert/strict";
import { fetchUserNameAsync } from "../src/async/await-patterns.js";
import { fetchUsersSequential, fetchUsersParallel } from "../src/async/flow.js";
import { fetchUsersInOrder } from "../src/async/iteration.js";
import { readTextFile } from "../src/async/promisify.js";

describe("fetchUserNameAsync", () => {
  it("returns user name", async () => {
    assert.equal(await fetchUserNameAsync("1"), "Ada");
  });

  it("throws wrapped error for missing user", async () => {
    await assert.rejects(() => fetchUserNameAsync("999"), /fetchUserName failed/);
  });
});

describe("sequential vs parallel", () => {
  it("both return same users for valid ids", async () => {
    const ids = ["1"];
    const seq = await fetchUsersSequential(ids);
    const par = await fetchUsersParallel(ids);
    assert.deepEqual(seq, par);
    assert.equal(seq[0]?.name, "Ada");
  });
});

describe("fetchUsersInOrder", () => {
  it("returns typed User array", async () => {
    const users = await fetchUsersInOrder(["1"]);
    assert.equal(users.length, 1);
    assert.equal(users[0]?.email, "ada@dev");
  });
});

describe("readTextFile", () => {
  it("reads package.json from project root", async () => {
    const text = await readTextFile("package.json");
    assert.match(text, /day-37-async-await/);
  });
});
