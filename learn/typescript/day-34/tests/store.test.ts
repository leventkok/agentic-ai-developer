import { describe, it } from "node:test";
import assert from "node:assert/strict";
import { join } from "node:path";
import { TaskStore } from "../src/core/store.js";

const testFileA = join(process.cwd(), ".test-tasks-a.json");
const testFileB = join(process.cwd(), ".test-tasks-b.json");

describe("TaskStore smoke tests", () => {
  it("creates and lists a task (happy path)", () => {
    const store = new TaskStore(testFileA);
    const created = store.create({ title: "Smoke test" });
    assert.equal(created.ok, true);
    if (created.ok) {
      assert.equal(store.list().length, 1);
      assert.equal(store.list()[0]?.title, "Smoke test");
    }
  });

  it("returns EMPTY_TITLE error (error path)", () => {
    const store = new TaskStore(testFileB);
    const result = store.create({ title: "  " });
    assert.equal(result.ok, false);
    if (!result.ok) {
      assert.equal(result.error.code, "EMPTY_TITLE");
    }
  });
});