import { mkdtemp, writeFile } from "node:fs/promises";
import os from "node:os";
import path from "node:path";
import { describe, it, after } from "node:test";
import assert from "node:assert/strict";
import { createNoteServer } from "../src/server/http-server.js";
import { NoteService } from "../src/services/note-service.js";
import { FileNoteStore } from "../src/storage/file-store.js";
import type { Server } from "node:http";

describe("Notes HTTP API", () => {
  let server: Server;
  let baseUrl = "";

  after(async () => {
    await new Promise<void>((resolve, reject) => server.close((err) => (err ? reject(err) : resolve())));
  });

  it("CRUD smoke test", async () => {
    const dir = await mkdtemp(path.join(os.tmpdir(), "notes-api-"));
    const file = path.join(dir, "notes.json");
    await writeFile(file, "[]");

    const service = new NoteService(new FileNoteStore(file));
    server = createNoteServer(
      { port: 0, host: "127.0.0.1", dataFile: file, env: "development" },
      service,
    );

    await new Promise<void>((resolve) => server.listen(0, "127.0.0.1", resolve));
    const addr = server.address();
    if (!addr || typeof addr === "string") throw new Error("no address");
    baseUrl = `http://127.0.0.1:${addr.port}`;

    const createRes = await fetch(`${baseUrl}/notes`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: "Test", body: "Body" }),
    });
    assert.equal(createRes.status, 201);
    const created = (await createRes.json()) as { id: string };

    const getRes = await fetch(`${baseUrl}/notes/${created.id}`);
    assert.equal(getRes.status, 200);

    const listRes = await fetch(`${baseUrl}/notes`);
    assert.equal(listRes.status, 200);

    const delRes = await fetch(`${baseUrl}/notes/${created.id}`, { method: "DELETE" });
    assert.equal(delRes.status, 204);
  });
});
