import type { Task } from "@core/types.js";



export function formatTask(task: Task, index: number): string {
    const mark = task.done ? "[x]" : "[ ]";
    return `${mark} ${index + 1}. ${task.title}`;
}