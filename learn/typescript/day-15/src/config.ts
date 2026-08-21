// Fake third-party library — augment in config.d.ts (Task 3)
export interface Config {
  apiUrl: string;
}

export function loadConfig(config: Config): void {
  console.log(`API: ${config.apiUrl}`);
}
