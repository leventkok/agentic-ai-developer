/** Typed event helpers — Day 42 */

export function onClick(element: Element, handler: (event: MouseEvent) => void): void {
  element.addEventListener("click", (event) => {
    if (!(event instanceof MouseEvent)) return;
    handler(event);
  });
}

export function onKeyDown(element: Element, handler: (event: KeyboardEvent) => void): void {
  element.addEventListener("keydown", (event) => {
    if (!(event instanceof KeyboardEvent)) return;
    handler(event);
  });
}

export function onSubmit(form: HTMLFormElement, handler: (event: SubmitEvent) => void): void {
  form.addEventListener("submit", (event) => {
    handler(event as SubmitEvent);
  });
}

export function delegateClick(
  parent: Element,
  selector: string,
  handler: (event: MouseEvent, target: HTMLElement) => void,
): void {
  parent.addEventListener("click", (event) => {
    if (!(event instanceof MouseEvent)) return;
    const target = event.target;
    if (!(target instanceof HTMLElement)) return;
    const match = target.closest(selector);
    if (!match || !parent.contains(match)) return;
    handler(event, match as HTMLElement);
  });
}

export type TodoAddedDetail = { id: string; title: string };

export const TODO_ADDED = "todo:added";

export function dispatchTodoAdded(root: EventTarget, detail: TodoAddedDetail): void {
  const doc = root instanceof Node ? root.ownerDocument : null;
  const CustomEventCtor = doc?.defaultView?.CustomEvent ?? CustomEvent;
  root.dispatchEvent(new CustomEventCtor(TODO_ADDED, { detail }));
}

export function onTodoAdded(
  root: EventTarget,
  handler: (detail: TodoAddedDetail) => void,
): void {
  root.addEventListener(TODO_ADDED, (event) => {
    handler((event as CustomEvent<TodoAddedDetail>).detail);
  });
}
