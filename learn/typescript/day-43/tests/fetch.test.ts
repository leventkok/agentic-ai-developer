import { describe, it, mock } from "node:test";
import assert from "node:assert/strict";
import { fetchJson, isWeatherData } from "../src/api/client.js";

describe("typed fetch", () => {
  it("parses successful JSON", async () => {
    const payload = { city: "Istanbul", tempC: 22, description: "Sunny" };
    globalThis.fetch = mock.fn(async () =>
      Response.json(payload),
    ) as typeof fetch;

    const result = await fetchJson<typeof payload>("https://example.com/weather");
    assert.equal(result.ok, true);
    if (result.ok) {
      assert.equal(result.data.city, "Istanbul");
    }
  });

  it("handles HTTP errors", async () => {
    globalThis.fetch = mock.fn(async () =>
      new Response("nope", { status: 500 }),
    ) as typeof fetch;

    const result = await fetchJson("https://example.com/weather");
    assert.equal(result.ok, false);
    if (!result.ok) {
      assert.equal(result.error.kind, "http");
    }
  });

  it("validates weather shape", () => {
    assert.equal(isWeatherData({ city: "X", tempC: 1, description: "Y" }), true);
    assert.equal(isWeatherData({ city: "X" }), false);
  });
});
