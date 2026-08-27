export type RetryOptions = {
  maxAttempts: number;
  baseDelayMs: number;
  isRetryable?: (err: unknown) => boolean;
};

function defaultIsRetryable(err: unknown): boolean {
  if (err instanceof Error) {
    return err.message.includes("transient");
  }
  return false;
}

function delay(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

/** Retry with exponential backoff — delay doubles each attempt */
export async function retryWithBackoff<T>(
  fn: () => Promise<T>,
  options: RetryOptions
): Promise<T> {
  const isRetryable = options.isRetryable ?? defaultIsRetryable;
  let lastError: unknown;

  for (let attempt = 1; attempt <= options.maxAttempts; attempt++) {
    try {
      return await fn();
    } catch (err: unknown) {
      lastError = err;
      if (attempt === options.maxAttempts || !isRetryable(err)) {
        throw err;
      }
      const backoff = options.baseDelayMs * 2 ** (attempt - 1);
      await delay(backoff);
    }
  }

  throw lastError;
}
