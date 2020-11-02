import * as readline from "readline"

const rl = readline.createInterface({ input: process.stdin, output: process.stdout })
const two = ["A", "B", "C"]
const three = ["D", "E", "F"]
const four = ["G", "H", "I"]
const five = ["J", "K", "L"]
const six = ["M", "N", "O"]
const seven = ["P", "Q", "R", "S"]
const eight = ["T", "U", "V"]
let lines: string[] = []
rl.on("line", line => lines.push(line)).on("close", () => {
    const string = lines[0]
    let time = 0

    for(let i=0; i<string.length; i++) {
        two.includes(string[i]) ? time += 3 : three.includes(string[i]) ? time += 4 : four.includes(string[i]) ? time += 5 : five.includes(string[i]) ? time += 6 : six.includes(string[i]) ? time += 7 : seven.includes(string[i]) ? time += 8 : eight.includes(string[i]) ? time += 9 : time += 10
    }

    console.log(time)
})