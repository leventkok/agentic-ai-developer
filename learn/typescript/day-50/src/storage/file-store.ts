import { mkdir, readFile, writeFile } from "node:fs/promises";
import path from "node:path";
import type { Note } from "../types/note.js";
import type { AppError } from "../types/errors.js";
import { err, ok, type Result } from "../types/result.js";

function isNodeError(error: unknown): error is NodeJS.ErrnoException {
  return error instanceof Error && "code" in error;
}

export class FileNoteStore {
  constructor(private readonly filePath: string) {}

  async load(): Promise<Result<Note[], AppError>> {
    try {
      const raw = await readFile(this.filePath, "utf8");
      const parsed = JSON.parse(raw) as unknown;
      if (!Array.isArray(parsed)) {
        return err({ kind: "storage", message: "Invalid notes file shape" });
      }
      return ok(parsed as Note[]);
    } catch (error) {
      if (isNodeError(error) && error.code === "ENOENT") {
        return ok([]);
      }
      const message = error instanceof Error ? error.message : "Read failed";
      return err({ kind: "storage", message });
    }
  }

  async save(notes: Note[]): Promise<Result<void, AppError>> {
    try {
      await mkdir(path.dirname(this.filePath), { recursive: true });
      await writeFile(this.filePath, JSON.stringify(notes, null, 2), "utf8");
      return ok(undefined);
    } catch (error) {
      const message = error instanceof Error ? error.message : "Write failed";
      return err({ kind: "storage", message });
    }
  }
}
