import { fetchUserNameAsync } from "./async/await-patterns.js";
import { fetchUsersSequential, fetchUsersParallel } from "./async/flow.js";
import { fetchUsersInOrder } from "./async/iteration.js";
import { readTextFile } from "./async/promisify.js";

export async function runDemo(): Promise<void> {
  console.log(await fetchUserNameAsync("1"));

  console.log("sequential:", await fetchUsersSequential(["1"]));
  console.log("parallel:", await fetchUsersParallel(["1"]));

  console.log("in order:", await fetchUsersInOrder(["1"]));

  try {
    const text = await readTextFile("package.json");
    console.log("read file chars:", text.length);
  } catch {
    console.log("readTextFile skipped");
  }
}

void runDemo();