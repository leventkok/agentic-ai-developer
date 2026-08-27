import { fetchUserWithResult, fetchUserOrThrow } from "./services/user-service.js";
import { formatApiError } from "./utils/safe-catch.js";
import { errorMessage } from "./utils/safe-catch.js";

export async function runDemo(): Promise<void> {
  const result = await fetchUserWithResult("1");
  if (result.ok) {
    console.log("Result ok:", result.value.name);
  } else {
    console.log("Result err:", formatApiError(result.error));
  }

  const missing = await fetchUserWithResult("999");
  if (!missing.ok) {
    console.log("Expected failure:", formatApiError(missing.error));
  }

  try {
    const user = await fetchUserOrThrow("1");
    console.log("Throw ok:", user.name);
  } catch (err: unknown) {
    console.log("Throw err:", errorMessage(err));
  }
}

void runDemo();