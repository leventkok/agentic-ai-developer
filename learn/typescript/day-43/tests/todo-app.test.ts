import { JSDOM } from "jsdom";
import { describe, it } from "node:test";
import assert from "node:assert/strict";
import { TodoApp, createTodoAppHtml } from "../src/app/todo-app.js";

describe("TodoApp", () => {
  it("adds toggles and deletes todos", () => {
    const dom = new JSDOM(createTodoAppHtml());
    const doc = dom.window.document;
    const app = new TodoApp({ root: doc.body });

    const input = doc.querySelector<HTMLInputElement>('input[name="title"]')!;
    input.value = "Learn TypeScript";
    doc.querySelector("form")!.dispatchEvent(
      new dom.window.Event("submit", { bubbles: true, cancelable: true }),
    );

    assert.equal(app.getTodos().length, 1);
    const id = app.getTodos()[0]!.id;

    app.toggle(id);
    assert.equal(app.getTodos()[0]!.done, true);

    app.delete(id);
    assert.equal(app.getTodos().length, 0);
  });
});
