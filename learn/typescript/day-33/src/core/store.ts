import { randomUUID } from "node:crypto";
import type { CreateTaskInput, Task, UpdateTaskInput } from "../types/task.js";
import type { Result } from "../types/result.js";
import { ok, err } from "../types/result.js";
import type { StoreError } from "./errors.js";

/**
 * In-memory task store with typed Result errors.
 *
 * @example
 * ```ts
 * const store = new TaskStore();
 * const result = store.create({ title: "Learn TypeScript" });
 * if (result.ok) console.log(result.value);
 * ```
 */
export class TaskStore {
    private tasks = new Map<string, Task>();

    /** Returns all tasks in insertion order. */
    list(): Task[] {
        return [...this.tasks.values()];
    }

    /**
     * Get a task by id.
     * @param id - UUID string
     * @returns `ok(value)` or `err(NOT_FOUND)`
     */
    get(id: string): Result<Task, StoreError> {
        const task = this.tasks.get(id);
        if (!task) {
            return err({ code: "NOT_FOUND", message: `Task with id ${id} not found` });
        }
        return ok(task);
    }

    /**
     * Create a new pending task.
     * @param input - `{ title: string }`
     * @returns `ok(task)` or `err(EMPTY_TITLE)`
     */
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
        return ok(task);
    }

    /**
     * Update an existing task (partial).
     * @returns `ok(task)`, `err(NOT_FOUND)`, or `err(EMPTY_TITLE)`
     */
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
        return ok(updated);
    }

    /**
     * Delete a task by id.
     * @returns `ok(undefined)` or `err(NOT_FOUND)`
     */
    delete(id: string): Result<void, StoreError> {
        if (!this.tasks.has(id)) {
            return err({ code: "NOT_FOUND", message: `Task with id ${id} not found` });
        }
        this.tasks.delete(id);
        return ok(undefined);
    }

    /** Mark a task as done. */
    complete(id: string): Result<Task, StoreError> {
        return this.update(id, { status: "done" });
    }
}
