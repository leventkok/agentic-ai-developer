// Module augmentation — extends ShopConfig from shop.ts (Task 3)
export {};

declare module "./shop" {
  interface ShopConfig {
    version?: string;
  }
}
