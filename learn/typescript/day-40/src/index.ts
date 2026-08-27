import { fetchUserWorkflow, resetFlakyAttempts, setFlakyFailures } from "./services/fetch-workflow.js";
import { debounceAsync } from "./async/debounce.js";
import { drainPending, getPendingCount, onGracefulShutdown } from "./async/shutdown.js";
import { formatApiError } from "./utils/safe-catch.js";

const debouncedFetch = debounceAsync(async (id: string) => fetchUserWorkflow(id), 50);

export async function runDemo(): Promise<void> {
  resetFlakyAttempts();
  setFlakyFailures(2);

  const result = await fetchUserWorkflow("1");
  if (result.ok) {
    console.log("Workflow success:", result.value.name);
  } else {
    console.log("Workflow error:", formatApiError(result.error));
  }

  debouncedFetch("1");
  debouncedFetch("1");
  const debounced = await debouncedFetch("1");
  if (debounced.ok) {
    console.log("Debounced:", debounced.value.name);
  }

  console.log("Pending before drain:", getPendingCount());
  await drainPending();
  console.log("Pending after drain:", getPendingCount());
}

onGracefulShutdown(async () => {
  await drainPending();
});

void runDemo();
