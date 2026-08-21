// Module augmentation — adds debug flag to Config (Task 3)
export {};

declare module "./config" {
  interface Config {
    debug?: boolean;
  }
}
