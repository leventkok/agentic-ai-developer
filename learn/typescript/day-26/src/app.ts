


import { greet } from "./utils/greet.js";
import type { User } from "./types/user.js";



const user: User = { id: 1, name: "John", email: "john@example.com" };

console.log(greet(user));
