import { JSDOM } from "jsdom";
import { describe, it } from "node:test";
import assert from "node:assert/strict";
import { onTodoAdded, TODO_ADDED } from "../src/dom/events.js";

describe("event typing", () => {
  it("custom event carries typed detail", () => {
    const dom = new JSDOM(`<div id="root"></div>`);
    const root = dom.window.document.getElementById("root")!;
    let received = "";

    onTodoAdded(root, (detail) => {
      received = detail.title;
    });

    root.dispatchEvent(
      new dom.window.CustomEvent(TODO_ADDED, { detail: { id: "1", title: "Test" } }),
    );

    assert.equal(received, "Test");
  });
});
