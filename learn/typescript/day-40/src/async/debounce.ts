/** Debounce async calls — only the last call within waitMs runs */
export function debounceAsync<TArgs extends unknown[], TResult>(
  fn: (...args: TArgs) => Promise<TResult>,
  waitMs: number
): (...args: TArgs) => Promise<TResult> {
  let timer: ReturnType<typeof setTimeout> | undefined;

  return (...args: TArgs): Promise<TResult> => {
    if (timer) clearTimeout(timer);

    return new Promise((resolve, reject) => {
      timer = setTimeout(() => {
        fn(...args).then(resolve).catch(reject);
      }, waitMs);
    });
  };
}
