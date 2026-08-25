

export type ParsedCommand = 
| { command: "list" }
| { command: "add"; title: string }
| { command: "done"; id: string }
| { command: "delete"; id: string }
| { command: "help" };


export function parseCommand(argv: string[]): ParsedCommand{
    const [, , cmd, ...rest] = argv;

    switch (cmd) {
        case "list":
          return { command: "list" };
        case "add": {
          const title = rest.join(" ").trim();
          return title ? { command: "add", title } : { command: "help" };
        }
        case "done": {
          const id = rest[0]?.trim();
          return id ? { command: "done", id } : { command: "help" };
        }
        case "delete": {
          const id = rest[0]?.trim();
          return id ? { command: "delete", id } : { command: "help" };
        }
        default:
          return { command: "help" };
    }
    
}