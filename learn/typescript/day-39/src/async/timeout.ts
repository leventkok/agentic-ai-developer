export function withTimeout<T>(
  work: (signal: AbortSignal) => Promise<T>,
  ms: number
): Promise<T> {
  const controller = new AbortController();
  const timer = setTimeout(() => controller.abort(), ms);

  return work(controller.signal).finally(() => clearTimeout(timer));
}
