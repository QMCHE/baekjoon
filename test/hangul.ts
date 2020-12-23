import * as readline from "readline"

const rl = readline.createInterface({ input: process.stdin, output: process.stdout })
const lines: string[] = []
rl.on("line", line => lines.push(line)).on("close", main)

function main() {
    const str = lines[0]
    for(let i=0; i<str.length; i++) {
        const unicode = str.charCodeAt(i).toString(16)
        console.log(unicode)
        console.log(String.fromCharCode(parseInt(unicode, 16)))
    }
}