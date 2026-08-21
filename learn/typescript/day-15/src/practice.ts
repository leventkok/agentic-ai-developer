// Day 15: Interfaces, Types & Classes — Practice

import { loadConfig } from "./config";

// Challenge 1: Plugin system
interface Plugin {
  name: string;
  version: string;
  execute(): void;
}

class LoyaltyPlugin implements Plugin {
  name = "loyalty";
  version = "1.0.0";

  execute(): void {
    console.log("Applying loyalty discount...");
  }
}

class MenuPrintPlugin implements Plugin {
  name = "menu-print";
  version = "2.0.0";

  execute(): void {
    console.log("Printing daily menu...");
  }
}

function runPlugins(plugins: Plugin[]): void {
  for (const plugin of plugins) {
    console.log(`Running ${plugin.name} v${plugin.version}`);
    plugin.execute();
  }
}

// Challenge 2: Repository contract
// Challenge 4: refactored MenuItem from interface to type alias
type MenuItem = {
  id: number;
  name: string;
  price: number;
};

// Trade-off: type alias fits plain object shapes; interface would allow
// declaration merging and cleaner extends — either works for MenuItem.

interface Repository<T> {
  findById(id: number): T | undefined;
  save(item: T): void;
}

class InMemoryRepository<T extends { id: number }> implements Repository<T> {
  private store = new Map<number, T>();

  findById(id: number): T | undefined {
    return this.store.get(id);
  }

  save(item: T): void {
    this.store.set(item.id, item);
  }
}

// Challenge 3: augmented config (debug added via config.d.ts)
const appConfig = {
  apiUrl: "https://api.masterfabric.cafe",
  debug: true,
};

// Run demos
runPlugins([new LoyaltyPlugin(), new MenuPrintPlugin()]);

const menuRepo = new InMemoryRepository<MenuItem>();
menuRepo.save({ id: 1, name: "Latte", price: 45 });
console.log("Found:", menuRepo.findById(1)?.name);

loadConfig(appConfig);

if (appConfig.debug) {
  console.log("Debug mode enabled");
}
