import * as readline from "readline"

const rl = readline.createInterface({ input: process.stdin, output: process.stdout })
let lines: string[] = []
rl.on("line", line => lines.push(line)).on("close", () => {
    let [a, b, c] = lines[0].split(" ").map(Number)
    
    let diff = c - b
    diff <= 0 ? console.log(-1) : console.log(Math.floor(a / diff) + 1)
})