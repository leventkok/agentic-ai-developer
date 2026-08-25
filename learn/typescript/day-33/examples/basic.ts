

import { TaskStore, formatTask} from "../src/index.js";

const store = new TaskStore();

const created = store.create({ title: "Learn TypeScript" });
if (!created.ok) {
    console.error(created.error);
    process.exit(1);
}

store.create({ title: "Build MVP" });
store.complete(created.value.id);
for (const task of store.list()) {
    console.log(formatTask(task));
}