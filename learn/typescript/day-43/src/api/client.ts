/** Typed fetch client — Day 44 */

import type { WeatherData } from "../types/todo.js";

export type ApiError =
  | { kind: "http"; status: number; message: string }
  | { kind: "network"; message: string }
  | { kind: "parse"; message: string };

export type FetchResult<T> =
  | { ok: true; data: T }
  | { ok: false; error: ApiError };

export interface FetchOptions {
  signal?: AbortSignal;
  method?: string;
  headers?: Record<string, string>;
  body?: unknown;
}

export async function fetchJson<T>(
  url: string,
  options: FetchOptions = {},
): Promise<FetchResult<T>> {
  try {
    const init: RequestInit = {
      signal: options.signal,
      method: options.method ?? "GET",
      headers: {
        Accept: "application/json",
        ...(options.body ? { "Content-Type": "application/json" } : {}),
        ...options.headers,
      },
      body: options.body ? JSON.stringify(options.body) : undefined,
    };

    const response = await fetch(url, init);
    if (!response.ok) {
      return {
        ok: false,
        error: {
          kind: "http",
          status: response.status,
          message: `HTTP ${response.status}`,
        },
      };
    }

    const data = (await response.json()) as T;
    return { ok: true, data };
  } catch (error) {
    if (error instanceof DOMException && error.name === "AbortError") {
      return { ok: false, error: { kind: "network", message: "Request aborted" } };
    }
    const message = error instanceof Error ? error.message : "Network error";
    return { ok: false, error: { kind: "network", message } };
  }
}

export function isWeatherData(value: unknown): value is WeatherData {
  if (typeof value !== "object" || value === null) return false;
  const v = value as Record<string, unknown>;
  return (
    typeof v.city === "string" &&
    typeof v.tempC === "number" &&
    typeof v.description === "string"
  );
}

export async function fetchWeather(
  url: string,
  signal?: AbortSignal,
): Promise<FetchResult<WeatherData>> {
  const result = await fetchJson<unknown>(url, { signal });
  if (!result.ok) return result;
  if (!isWeatherData(result.data)) {
    return { ok: false, error: { kind: "parse", message: "Invalid weather payload" } };
  }
  return { ok: true, data: result.data };
}
