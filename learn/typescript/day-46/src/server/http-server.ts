import { createServer, type IncomingMessage, type Server, type ServerResponse } from "node:http";
import type { AppConfig } from "../config/env.js";
import { handleNotesRoute } from "../handlers/notes.js";
import type { NoteService } from "../services/note-service.js";

export function createNoteServer(config: AppConfig, service: NoteService): Server {
  return createServer(async (req, res) => {
    try {
      await route(req, res, service);
    } catch {
      res.writeHead(500, { "Content-Type": "application/json" });
      res.end(JSON.stringify({ error: { kind: "storage", message: "Internal error" } }));
    }
  });
}

async function route(req: IncomingMessage, res: ServerResponse, service: NoteService): Promise<void> {
  const url = new URL(req.url ?? "/", "http://localhost");
  if (url.pathname === "/health") {
    res.writeHead(200, { "Content-Type": "application/json" });
    res.end(JSON.stringify({ status: "ok" }));
    return;
  }

  const match = url.pathname.match(/^\/notes(?:\/([^/]+))?$/);
  if (!match) {
    res.writeHead(404, { "Content-Type": "application/json" });
    res.end(JSON.stringify({ error: { kind: "not_found", id: url.pathname } }));
    return;
  }

  await handleNotesRoute(req, res, service, match[1]);
}
