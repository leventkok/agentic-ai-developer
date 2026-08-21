interface User {
    id: number;
    name: string;
    email: string;
  }
  
  const user: User = {
    id: 1,
    name: "Ali",
    email: "ali@masterfabric.com",
  };
  
  type Customer = {
    id: number;
    name: string;
    email: string;
  };
  
  interface AdminUser extends User {
    role: "admin";
    permissions: string[];
  }
  
  type Barista = Customer & {
    role: "barista";
    shift: "morning" | "evening";
  };
  
  const admin: AdminUser = {
    id: 2,
    name: "Admin",
    email: "admin@masterfabric.com",
    role: "admin",
    permissions: ["manage_menu"],
  };
  
  const barista: Barista = {
    id: 3,
    name: "Zeynep",
    email: "zeynep@masterfabric.com",
    role: "barista",
    shift: "morning",
  };
  
  function printUser(u: User): void {
    console.log(`${u.name} <${u.email}>`);
  }

  // interface: object shapes you extend (AdminUser extends User)
  // type: unions, intersections, and aliases (Barista = Customer & {...})

  // Run demos
  printUser(user);
  printUser(admin);
  console.log("Barista shift:", barista.shift);