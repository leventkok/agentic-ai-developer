


import type { PublicTask, Task } from "../types/task.js";


export function formatTask(task: Task) : string {
    const mark = task.status === "done" ? "✅" : "⏳";
    return `${mark} ${task.title} (${task.id.slice(0, 8)}...)`;
}

export function toPublicTask(task: Task) : PublicTask {
    return{
        id: task.id,
        title: task.title,
        status: task.status,
        createdAt: task.createdAt.toISOString(),
        updatedAt: task.updatedAt.toISOString(),
    };
}