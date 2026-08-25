import { existsSync, readFileSync, writeFileSync } from "node:fs";
import { join } from "node:path";
import { randomUUID } from "node:crypto";
import type { CreateTaskInput, Task, UpdateTaskInput } from "../types/task.js";
import type { Result } from "../types/result.js";
import { ok, err } from "../types/result.js";
import type { StoreError } from "./errors.js";

const DEFAULT_DATA_FILE = join(process.cwd(), "tasks.json");

type StoredTask = Omit<Task, "createdAt" | "updatedAt"> & {
    createdAt: string;
    updatedAt: string;
};

/**
 * In-memory task store with typed Result errors.
 * Persists to tasks.json so CLI commands share state across runs.
 */
export class TaskStore {
    private tasks = new Map<string, Task>();
    private readonly dataFile: string;

    constructor(dataFile: string = DEFAULT_DATA_FILE) {
        this.dataFile = dataFile;
        this.load();
    }

    private load(): void {
        if (!existsSync(this.dataFile)) return;
        const raw = JSON.parse(readFileSync(this.dataFile, "utf-8")) as StoredTask[];
        for (const item of raw) {
            this.tasks.set(item.id, {
                ...item,
                createdAt: new Date(item.createdAt),
                updatedAt: new Date(item.updatedAt),
            });
        }
    }

    private save(): void {
        const data: StoredTask[] = [...this.tasks.values()].map((task) => ({
            id: task.id,
            title: task.title,
            status: task.status,
            createdAt: task.createdAt.toISOString(),
            updatedAt: task.updatedAt.toISOString(),
        }));
        writeFileSync(this.dataFile, JSON.stringify(data, null, 2));
    }

    list(): Task[] {
        return [...this.tasks.values()];
    }

    get(id: string): Result<Task, StoreError> {
        const task = this.tasks.get(id);
        if (!task) {
            return err({ code: "NOT_FOUND", message: `Task with id ${id} not found` });
        }
        return ok(task);
    }

    create(input: CreateTaskInput): Result<Task, StoreError> {
        const title = input.title.trim();
        if (!title) {
            return err({ code: "EMPTY_TITLE", message: "Title is required" });
        }

        const now = new Date();
        const task: Task = {
            id: randomUUID(),
            title,
            status: "pending",
            createdAt: now,
            updatedAt: now,
        };
        this.tasks.set(task.id, task);
        this.save();
        return ok(task);
    }

    update(id: string, input: UpdateTaskInput): Result<Task, StoreError> {
        const existing = this.tasks.get(id);
        if (!existing) {
            return err({ code: "NOT_FOUND", message: `Task with id ${id} not found` });
        }
        const updated: Task = {
            ...existing,
            title: input.title?.trim() ?? existing.title,
            status: input.status ?? existing.status,
            updatedAt: new Date(),
        };
        if (!updated.title) {
            return err({ code: "EMPTY_TITLE", message: "Title is required" });
        }
        this.tasks.set(id, updated);
        this.save();
        return ok(updated);
    }

    delete(id: string): Result<void, StoreError> {
        if (!this.tasks.has(id)) {
            return err({ code: "NOT_FOUND", message: `Task with id ${id} not found` });
        }
        this.tasks.delete(id);
        this.save();
        return ok(undefined);
    }

    complete(id: string): Result<Task, StoreError> {
        return this.update(id, { status: "done" });
    }
}
