// Simulates a small library module we will augment in shop.d.ts
export interface ShopConfig {
  shopName: string;
}

export function initShop(config: ShopConfig): void {
  console.log(`Shop: ${config.shopName}`);
}

