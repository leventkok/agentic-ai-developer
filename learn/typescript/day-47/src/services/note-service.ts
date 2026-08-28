import { randomUUID } from "node:crypto";
import type { AppError } from "../types/errors.js";
import type { CreateNoteInput, Note, UpdateNoteInput } from "../types/note.js";
import { err, ok, type Result } from "../types/result.js";
import { FileNoteStore } from "../storage/file-store.js";

export class NoteService {
  constructor(private readonly store: FileNoteStore) {}

  async list(): Promise<Result<Note[], AppError>> {
    return this.store.load();
  }

  async get(id: string): Promise<Result<Note, AppError>> {
    const loaded = await this.store.load();
    if (!loaded.ok) return loaded;
    const note = loaded.value.find((n) => n.id === id);
    if (!note) return err({ kind: "not_found", id });
    return ok(note);
  }

  async create(input: CreateNoteInput): Promise<Result<Note, AppError>> {
    const validation = validateCreate(input);
    if (validation) return err(validation);

    const loaded = await this.store.load();
    if (!loaded.ok) return loaded;

    const note: Note = {
      id: randomUUID(),
      title: input.title.trim(),
      body: input.body.trim(),
      createdAt: new Date().toISOString(),
    };

    loaded.value.push(note);
    const saved = await this.store.save(loaded.value);
    if (!saved.ok) return saved;
    return ok(note);
  }

  async update(id: string, input: UpdateNoteInput): Promise<Result<Note, AppError>> {
    const loaded = await this.store.load();
    if (!loaded.ok) return loaded;

    const index = loaded.value.findIndex((n) => n.id === id);
    if (index === -1) return err({ kind: "not_found", id });

    const current = loaded.value[index]!;
    if (input.title !== undefined) {
      const title = input.title.trim();
      if (!title) return err({ kind: "validation", field: "title", message: "Title required" });
      current.title = title;
    }
    if (input.body !== undefined) {
      current.body = input.body.trim();
    }

    loaded.value[index] = current;
    const saved = await this.store.save(loaded.value);
    if (!saved.ok) return saved;
    return ok(current);
  }

  async delete(id: string): Promise<Result<void, AppError>> {
    const loaded = await this.store.load();
    if (!loaded.ok) return loaded;

    const next = loaded.value.filter((n) => n.id !== id);
    if (next.length === loaded.value.length) {
      return err({ kind: "not_found", id });
    }

    return this.store.save(next);
  }
}

function validateCreate(input: CreateNoteInput): AppError | null {
  if (!input.title?.trim()) {
    return { kind: "validation", field: "title", message: "Title is required" };
  }
  return null;
}
