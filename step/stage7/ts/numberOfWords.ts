import * as readline from "readline"

const rl = readline.createInterface({ input: process.stdin, output: process.stdout })
let lines: string[] = []
rl.on("line", line => lines.push(line)).on("close", solution)

function solution() {
    console.log(lines[0].split(" ").filter(a => a !== "").length)
}