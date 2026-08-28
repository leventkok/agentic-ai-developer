import type { IncomingMessage, ServerResponse } from "node:http";
import type { NoteService } from "../services/note-service.js";
import { errorResponse } from "../types/errors.js";
import type { CreateNoteInput, UpdateNoteInput } from "../types/note.js";

export async function handleNotesRoute(
  req: IncomingMessage,
  res: ServerResponse,
  service: NoteService,
  noteId?: string,
): Promise<void> {
  const method = req.method ?? "GET";

  if (method === "GET" && !noteId) {
    const result = await service.list();
    if (!result.ok) return sendError(res, result.error);
    return sendJson(res, 200, result.value);
  }

  if (method === "GET" && noteId) {
    const result = await service.get(noteId);
    if (!result.ok) return sendError(res, result.error);
    return sendJson(res, 200, result.value);
  }

  if (method === "POST" && !noteId) {
    const body = await readJsonBody<CreateNoteInput>(req);
    if (!body.ok) return sendJson(res, 400, { error: body.error });
    const result = await service.create(body.value);
    if (!result.ok) return sendError(res, result.error);
    return sendJson(res, 201, result.value);
  }

  if (method === "PATCH" && noteId) {
    const body = await readJsonBody<UpdateNoteInput>(req);
    if (!body.ok) return sendJson(res, 400, { error: body.error });
    const result = await service.update(noteId, body.value);
    if (!result.ok) return sendError(res, result.error);
    return sendJson(res, 200, result.value);
  }

  if (method === "DELETE" && noteId) {
    const result = await service.delete(noteId);
    if (!result.ok) return sendError(res, result.error);
    res.writeHead(204);
    res.end();
    return;
  }

  sendJson(res, 405, { error: { kind: "validation", field: "method", message: "Method not allowed" } });
}

function sendJson(res: ServerResponse, status: number, body: unknown): void {
  res.writeHead(status, { "Content-Type": "application/json" });
  res.end(JSON.stringify(body));
}

function sendError(res: ServerResponse, error: Parameters<typeof errorResponse>[0]): void {
  const mapped = errorResponse(error);
  sendJson(res, mapped.status, mapped.body);
}

async function readJsonBody<T>(req: IncomingMessage): Promise<
  | { ok: true; value: T }
  | { ok: false; error: { kind: "validation"; field: string; message: string } }
> {
  const chunks: Buffer[] = [];
  for await (const chunk of req) {
    chunks.push(typeof chunk === "string" ? Buffer.from(chunk) : chunk);
  }
  if (chunks.length === 0) {
    return { ok: false, error: { kind: "validation", field: "body", message: "JSON body required" } };
  }
  try {
    return { ok: true, value: JSON.parse(Buffer.concat(chunks).toString("utf8")) as T };
  } catch {
    return { ok: false, error: { kind: "validation", field: "body", message: "Invalid JSON" } };
  }
}
