import { existsSync, readFileSync, writeFileSync } from "fs";
import { join } from "path";
import type { Task } from "@core/types.js";

const DATA_FILE = join(process.cwd(), "tasks.json");

type StoreData = {
    tasks: Task[];
    nextId: number;
};

export class TaskStore {
    private tasks: Task[] = [];
    private nextId = 1;

    constructor() {
        this.load();
    }

    private load(): void {
        if (!existsSync(DATA_FILE)) return;
        const data = JSON.parse(readFileSync(DATA_FILE, "utf-8")) as StoreData;
        this.tasks = data.tasks;
        this.nextId = data.nextId;
    }

    private save(): void {
        const data: StoreData = { tasks: this.tasks, nextId: this.nextId };
        writeFileSync(DATA_FILE, JSON.stringify(data, null, 2));
    }

    add(title: string): Task {
        const task: Task = { id: this.nextId++, title, done: false };
        this.tasks.push(task);
        this.save();
        return task;
    }

    list(): Task[] {
        return [...this.tasks];
    }

    getByIndex(index: number): Task | undefined {
        return this.tasks[index];
    }

    toggle(index: number): Task | undefined {
        const task = this.getByIndex(index);
        if (!task) return undefined;
        task.done = !task.done;
        this.save();
        return task;
    }
}
