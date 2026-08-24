import { add } from "@shared/math.js";
import { formatResult } from "@/lib/format.js";

const sum = add(10, 5);
console.log(formatResult("sum", sum));
