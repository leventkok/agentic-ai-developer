

import { getNameFromArgs } from "./parseArgs";


function greet(name: string): string {
    return `Hello, ${name}!`;
}



function formatToday(): string{
    const today = new Date();

    return today.toLocaleDateString("en-US", {
        weekday: "long",
        year: "numeric",
        month: "long",
        day: "numeric",
    });
}


const name = getNameFromArgs(process.argv);

if (name === undefined) {
    console.error("Error: Please provide a name.");
    console.error("Usage: node dist/greet.js <name>");
    process.exit(1);
}

console.log(`${greet(name)} Today is ${formatToday()}`);
