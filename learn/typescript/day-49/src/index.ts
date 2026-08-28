import { loadConfig } from "./config/env.js";
import { createNoteServer } from "./server/http-server.js";
import { NoteService } from "./services/note-service.js";
import { FileNoteStore } from "./storage/file-store.js";

const config = loadConfig();
if ("kind" in config) {
  console.error("Config error:", config.message);
  process.exit(1);
}

const store = new FileNoteStore(config.dataFile);
const service = new NoteService(store);
const server = createNoteServer(config, service);

server.listen(config.port, config.host, () => {
  console.log(`Notes API [${config.env}] http://${config.host}:${config.port}`);
});

process.on("SIGINT", () => {
  server.close(() => process.exit(0));
});
