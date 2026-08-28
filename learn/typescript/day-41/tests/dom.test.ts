import { JSDOM } from "jsdom";
import { describe, it } from "node:test";
import assert from "node:assert/strict";
import { requireInput } from "../src/dom/query.js";

describe("DOM query helpers", () => {
  it("requireInput returns typed input element", () => {
    const dom = new JSDOM(`<form><input name="title" type="text" /></form>`);
    const input = requireInput(dom.window.document, "input");
    assert.equal(input.name, "title");
  });

  it("throws when element missing", () => {
    const dom = new JSDOM(`<div></div>`);
    assert.throws(() => requireInput(dom.window.document, "input"));
  });
});
