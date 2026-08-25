
import { TaskStore } from "../src/index.js";

const store = new TaskStore();

const empty = store.create({ title: "   " });
console.log("Empty title:", empty);

const missing = store.get("not-a-real-id");
console.log("Not found:", missing);