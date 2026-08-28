import { JSDOM } from "jsdom";
import { describe, it } from "node:test";
import assert from "node:assert/strict";
import { readTodoForm, validateTodoForm } from "../src/dom/forms.js";

describe("form handling", () => {
  it("reads and validates todo form", () => {
    const dom = new JSDOM(`
      <form id="f">
        <input name="title" value="  Buy milk  " />
      </form>
    `);
    const form = dom.window.document.getElementById("f") as HTMLFormElement;
    const state = readTodoForm(form);
    assert.equal(state.title, "Buy milk");
    assert.equal(validateTodoForm(state), null);
  });

  it("rejects empty title", () => {
    assert.equal(validateTodoForm({ title: "" }), "Title is required");
  });
});
