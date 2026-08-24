import type { User } from "../types/user.js";

export function greet(user: User): string {
    return `Hello, ${user.name}!`;
}
