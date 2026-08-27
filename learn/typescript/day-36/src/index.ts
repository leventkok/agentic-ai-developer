import { fetchUser, logUser, fetchUserName } from "./promises/basic.js";
import { fetchUserAndPosts, fetchAllSettled } from "./promises/parallel.js";

export async function runDemo(): Promise<void> {
  const user = await fetchUser("1");
  console.log(user);

  await logUser("1");

  const name = await fetchUserName("1");
  console.log(name);

  const [u, posts] = await fetchUserAndPosts("1");
  console.log(u.name, posts);

  const results = await fetchAllSettled(["1", "999"]);
  console.log(results);
}

void runDemo();