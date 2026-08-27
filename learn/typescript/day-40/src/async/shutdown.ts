const pending = new Set<Promise<unknown>>();

export function trackPromise<T>(promise: Promise<T>): Promise<T> {
  pending.add(promise);
  return promise.finally(() => pending.delete(promise));
}

export function getPendingCount(): number {
  return pending.size;
}

export async function drainPending(): Promise<void> {
  await Promise.allSettled([...pending]);
}

export function onGracefulShutdown(handler: () => Promise<void>): void {
  const shutdown = async () => {
    console.log("SIGINT received — draining pending work...");
    await handler();
    console.log("Shutdown complete");
    process.exit(0);
  };

  process.once("SIGINT", () => {
    void shutdown();
  });
}
