/** Typed form handling — Day 43 */

import type { TodoFormState } from "../types/todo.js";

export function readTodoForm(form: HTMLFormElement): TodoFormState {
  const titleInput = form.elements.namedItem("title");
  if (!titleInput || titleInput.tagName !== "INPUT") {
    throw new Error("title field is required");
  }
  const value = (titleInput as HTMLInputElement).value;
  return { title: value.trim() };
}

export function validateTodoForm(state: TodoFormState): string | null {
  if (!state.title) {
    return "Title is required";
  }
  if (state.title.length > 100) {
    return "Title is too long";
  }
  return null;
}

export function syncFormState(form: HTMLFormElement, state: TodoFormState): void {
  const input = form.elements.namedItem("title");
  if (input && input.tagName === "INPUT") {
    (input as HTMLInputElement).value = state.title;
  }
}

/** Read all fields via FormData when running in a real browser. */
export function readTodoFormData(form: HTMLFormElement): TodoFormState {
  const data = new FormData(form);
  const title = data.get("title");
  if (typeof title !== "string") {
    throw new Error("title field is required");
  }
  return { title: title.trim() };
}
