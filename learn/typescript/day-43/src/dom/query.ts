/** Typed DOM query helpers — Day 41 */

export function requireElement<T extends Element>(
  root: ParentNode,
  selector: string,
): T {
  const el = root.querySelector<T>(selector);
  if (!el) {
    throw new Error(`Element not found: ${selector}`);
  }
  return el;
}

export function requireInput(root: ParentNode, selector: string): HTMLInputElement {
  return requireElement<HTMLInputElement>(root, selector);
}

export function requireForm(root: ParentNode, selector: string): HTMLFormElement {
  return requireElement<HTMLFormElement>(root, selector);
}

export function requireList(root: ParentNode, selector: string): HTMLUListElement {
  return requireElement<HTMLUListElement>(root, selector);
}
