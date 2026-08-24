import { TaskStore } from "@core/store.js";
import { formatTask } from "@utils/format.js";

const store = new TaskStore();
const [, , command, ...rest] = process.argv;

if (command === "list") {
    store.list().forEach((task, i) => console.log(formatTask(task, i)));
} else if (command === "add") {
    const title = rest.join(" ");
    if (!title) { console.error("Usage: add <title>"); process.exit(1); }
    store.add(title);
    console.log(`Added: ${title}`);
} else if (command === "toggle") {
    const index = Number(rest[0]);
    const task = store.toggle(index);
    if (task === undefined) {
        console.error("Task not found");
        process.exit(1);
    }
    console.log(formatTask(task, index));
} else {
    console.log("Commands: list | add <title> | toggle <index>");
}