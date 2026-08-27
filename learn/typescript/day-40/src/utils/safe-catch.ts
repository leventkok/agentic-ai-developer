export function errorMessage(err: unknown): string {
  if (err instanceof Error) {
    return err.message;
  }
  if (typeof err === "string") {
    return err;
  }
  return "Unknown error";
}

export function formatApiError(error: import("../types/errors.js").ApiError): string {
  switch (error.kind) {
    case "network":
      return `Network error: ${error.message}`;
    case "validation":
      return `Validation error on ${error.field}: ${error.message}`;
    case "not_found":
      return `${error.resource} ${error.id} not found`;
  }
}