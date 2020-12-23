import * as readline from "readline"

const rl = readline.createInterface({ input: process.stdin, output: process.stdout })
let lines: string[] = []
rl.on("line", line => lines.push(line)).on("close", () => {
    const n = Number(lines[0]) - 1

    console.log()
})