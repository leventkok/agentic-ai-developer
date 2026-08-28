import { describe, it } from "node:test";
import assert from "node:assert/strict";
import { loadConfig } from "../src/config/env.js";

describe("loadConfig", () => {
  it("parses valid env", () => {
    const cfg = loadConfig({
      PORT: "4000",
      HOST: "127.0.0.1",
      ENV: "development",
      DATA_FILE: "./data/notes.json",
    });
    assert.equal("port" in cfg && cfg.port, 4000);
  });

  it("fails when PORT missing", () => {
    const cfg = loadConfig({ DATA_FILE: "./data/notes.json" });
    assert.equal("kind" in cfg && cfg.kind, "config");
  });
});
