import { readFile } from "node:fs";

export function readTextFile(path: string): Promise<string> {
  return new Promise((resolve, reject) => {
    readFile(path, "utf-8", (err, data) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(data);
    });
  });
}