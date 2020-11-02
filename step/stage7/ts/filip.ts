import * as readline from "readline"

const rl = readline.createInterface({ input: process.stdin, output: process.stdout })
let lines: string[] = []
rl.on("line", line => lines.push(line)).on("close", () => {
    let [a, b] = lines[0].split(" ")
    a = a.toString().split("").reverse().map(Number).join("")
    b = b.toString().split("").reverse().map(Number).join("")
    a > b ? console.log(a) : console.log(b)
})