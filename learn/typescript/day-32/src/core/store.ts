

import { randomUUID } from "node:crypto";
import type { CreateTaskInput, Task, UpdateTaskInput } from "../types/task.js";
import type { Result } from "../types/result.js";
import { ok, err } from "../types/result.js";
import type { StoreError } from "./errors.js";

export class TaskStore {
    private tasks = new Map<String, Task>();

    list() : Task[]{
        return [...this.tasks.values()];
    }

    get(id: string) : Result<Task, StoreError> {
        const task = this.tasks.get(id);
        if (!task) {
            return err({ code: "NOT_FOUND", message: `Task with id ${id} not found` });
        }
        return ok(task);
    }

    create(input: CreateTaskInput) : Result<Task, StoreError> {
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

    update(id: string, input: UpdateTaskInput) : Result<Task, StoreError> {
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
        if(!updated.title) {
            return err({ code: "EMPTY_TITLE", message: "Title is required" });
        }
        this.tasks.set(id, updated);
        return ok(updated);
    }

    delete(id: string) : Result<void, StoreError> {
        if (!this.tasks.has(id)) {
            return err({ code: "NOT_FOUND", message: `Task with id ${id} not found` });

        }
        this.tasks.delete(id);
        return ok(undefined);
    }

    complete(id: string) : Result<Task, StoreError> {
        return this.update(id, { status: "done" });
    }
}