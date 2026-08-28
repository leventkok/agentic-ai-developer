/** Typed todo app — Day 45 capstone */

import { fetchWeather } from "../api/client.js";
import { delegateClick, dispatchTodoAdded, onSubmit } from "../dom/events.js";
import { readTodoForm, validateTodoForm } from "../dom/forms.js";
import { requireElement, requireForm, requireList } from "../dom/query.js";
import type { Todo, WeatherData } from "../types/todo.js";

export interface TodoAppOptions {
  root: ParentNode;
  weatherUrl?: string;
}

export class TodoApp {
  private readonly form: HTMLFormElement;
  private readonly list: HTMLUListElement;
  private readonly errorEl: HTMLElement;
  private readonly weatherEl: HTMLElement;
  private readonly todos: Todo[] = [];
  private readonly weatherUrl?: string;

  constructor(options: TodoAppOptions) {
    this.form = requireForm(options.root, "#todo-form");
    this.list = requireList(options.root, "#todo-list");
    this.errorEl = requireElement<HTMLElement>(options.root, "#todo-error");
    this.weatherUrl = options.weatherUrl;
    this.weatherEl = requireElement<HTMLElement>(options.root, "#weather");

    this.bindEvents();
  }

  private bindEvents(): void {
    onSubmit(this.form, (event) => {
      event.preventDefault();
      this.handleAdd();
    });

    delegateClick(this.list, "[data-action='toggle']", (_event, target) => {
      const id = target.dataset.id;
      if (!id) return;
      this.toggle(id);
    });

    delegateClick(this.list, "[data-action='delete']", (_event, target) => {
      const id = target.dataset.id;
      if (!id) return;
      this.delete(id);
    });
  }

  private handleAdd(): void {
    let state;
    try {
      state = readTodoForm(this.form);
    } catch {
      this.showError("Invalid form data");
      return;
    }

    const validationError = validateTodoForm(state);
    if (validationError) {
      this.showError(validationError);
      return;
    }

    const todo: Todo = {
      id: crypto.randomUUID(),
      title: state.title,
      done: false,
    };
    this.todos.push(todo);
    this.form.reset();
    this.showError("");
    this.render();
    dispatchTodoAdded(this.form, { id: todo.id, title: todo.title });
  }

  toggle(id: string): void {
    const todo = this.todos.find((t) => t.id === id);
    if (!todo) return;
    todo.done = !todo.done;
    this.render();
  }

  delete(id: string): void {
    const index = this.todos.findIndex((t) => t.id === id);
    if (index === -1) return;
    this.todos.splice(index, 1);
    this.render();
  }

  async loadWeather(signal?: AbortSignal): Promise<WeatherData | null> {
    if (!this.weatherUrl) return null;
    const result = await fetchWeather(this.weatherUrl, signal);
    if (!result.ok) {
      this.weatherEl.textContent = result.error.message;
      return null;
    }
    this.weatherEl.textContent = `${result.data.city}: ${result.data.tempC}°C ${result.data.description}`;
    return result.data;
  }

  getTodos(): readonly Todo[] {
    return this.todos;
  }

  private showError(message: string): void {
    this.errorEl.textContent = message;
  }

  private render(): void {
    const doc = this.list.ownerDocument;
    this.list.replaceChildren();
    for (const todo of this.todos) {
      const li = doc.createElement("li");
      li.dataset.id = todo.id;

      const span = doc.createElement("span");
      span.textContent = todo.title;
      if (todo.done) span.style.textDecoration = "line-through";

      const toggleBtn = doc.createElement("button");
      toggleBtn.type = "button";
      toggleBtn.dataset.action = "toggle";
      toggleBtn.dataset.id = todo.id;
      toggleBtn.textContent = todo.done ? "Undo" : "Done";

      const deleteBtn = doc.createElement("button");
      deleteBtn.type = "button";
      deleteBtn.dataset.action = "delete";
      deleteBtn.dataset.id = todo.id;
      deleteBtn.textContent = "Delete";

      li.append(span, toggleBtn, deleteBtn);
      this.list.append(li);
    }
  }
}

export function createTodoAppHtml(): string {
  return `
    <div id="app">
      <form id="todo-form">
        <input name="title" type="text" placeholder="New todo" />
        <button type="submit">Add</button>
      </form>
      <p id="todo-error"></p>
      <ul id="todo-list"></ul>
      <div id="weather"></div>
    </div>
  `;
}
