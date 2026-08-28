import { mkdtemp, writeFile } from "node:fs/promises";
import os from "node:os";
import path from "node:path";
import { describe, it } from "node:test";
import assert from "node:assert/strict";
import { FileNoteStore } from "../src/storage/file-store.js";
import { NoteService } from "../src/services/note-service.js";

describe("FileNoteStore", () => {
  it("loads empty array when file missing", async () => {
    const dir = await mkdtemp(path.join(os.tmpdir(), "notes-"));
    const store = new FileNoteStore(path.join(dir, "notes.json"));
    const result = await store.load();
    assert.equal(result.ok, true);
    if (result.ok) assert.deepEqual(result.value, []);
  });

  it("creates and lists notes", async () => {
    const dir = await mkdtemp(path.join(os.tmpdir(), "notes-"));
    const file = path.join(dir, "notes.json");
    await writeFile(file, "[]");
    const service = new NoteService(new FileNoteStore(file));

    const created = await service.create({ title: "Hello", body: "World" });
    assert.equal(created.ok, true);

    const list = await service.list();
    assert.equal(list.ok, true);
    if (list.ok) assert.equal(list.value.length, 1);
  });
});
