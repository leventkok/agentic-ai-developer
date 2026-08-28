export type ValidationError = { kind: "validation"; field: string; message: string };
export type NotFoundError = { kind: "not_found"; id: string };
export type StorageError = { kind: "storage"; message: string };
export type ConfigError = { kind: "config"; message: string };

export type AppError = ValidationError | NotFoundError | StorageError | ConfigError;

export function errorResponse(error: AppError): { status: number; body: { error: AppError } } {
  switch (error.kind) {
    case "validation":
      return { status: 400, body: { error } };
    case "not_found":
      return { status: 404, body: { error } };
    case "config":
      return { status: 500, body: { error } };
    case "storage":
      return { status: 500, body: { error } };
  }
}
