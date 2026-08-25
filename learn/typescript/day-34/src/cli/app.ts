import { TaskStore } from "../core/store.js";
import { formatTask } from "../utils/format.js";
import { parseCommand } from "../utils/parse.js";

const store = new TaskStore();
const parsed = parseCommand(process.argv);

switch (parsed.command) {
  case "list":
    for (const task of store.list()) {
      console.log(formatTask(task));
    }
    break;

  case "add": {
    const result = store.create({ title: parsed.title });
    if (!result.ok) {
      console.error(result.error.message);
      process.exit(1);
    }
    console.log("Added:", formatTask(result.value));
    break;
  }

  case "done": {
    const result = store.complete(parsed.id);
    if (!result.ok) {
      console.error(result.error.message);
      process.exit(1);
    }
    console.log("Completed:", formatTask(result.value));
    break;
  }

  case "delete": {
    const result = store.delete(parsed.id);
    if (!result.ok) {
      console.error(result.error.message);
      process.exit(1);
    }
    console.log("Deleted.");
    break;
  }

  case "help":
  default:
    console.log("Commands: list | add <title> | done <id> | delete <id>");
}
