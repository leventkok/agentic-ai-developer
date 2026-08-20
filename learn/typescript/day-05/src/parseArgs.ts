export function getNameFromArgs(args: string[]): string | undefined {
  const name = args[2];

  if (name === undefined) {
    return undefined;
  }

  const trimmed = name.trim();

  if (trimmed === "") {
    return undefined;
  }

  return trimmed;
}
