import type { ConfigError } from "../types/errors.js";

export interface AppConfig {
  port: number;
  host: string;
  dataFile: string;
  env: "development" | "staging" | "production";
}

function parsePort(raw: string | undefined): number | ConfigError {
  if (!raw) return { kind: "config", message: "PORT is required" };
  const port = Number(raw);
  if (!Number.isInteger(port) || port <= 0) {
    return { kind: "config", message: `Invalid PORT: ${raw}` };
  }
  return port;
}

function parseEnv(raw: string | undefined): AppConfig["env"] | ConfigError {
  const value = raw ?? "development";
  if (value === "development" || value === "staging" || value === "production") {
    return value;
  }
  return { kind: "config", message: `Invalid ENV: ${value}` };
}

export function loadConfig(env: NodeJS.ProcessEnv = process.env): AppConfig | ConfigError {
  const port = parsePort(env.PORT);
  if (typeof port !== "number") return port;

  const appEnv = parseEnv(env.ENV);
  if (typeof appEnv !== "string") return appEnv;

  const dataFile = env.DATA_FILE?.trim();
  if (!dataFile) {
    return { kind: "config", message: "DATA_FILE is required" };
  }

  return {
    port,
    host: env.HOST?.trim() || "127.0.0.1",
    dataFile,
    env: appEnv,
  };
}
