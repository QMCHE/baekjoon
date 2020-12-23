import * as readline from "readline"

const rl = readline.createInterface({ input: process.stdin, output: process.stdout })
let lines: string[] = []
rl.on("line", (line: string) => lines.push(line)).on("close", () => {
    let weight = Number(lines[0])
    let fiveCount = Math.floor(weight/5)
    let threeCount = Math.floor((weight-(fiveCount*5))/3)

    while(true) {
        if(weight === fiveCount * 5 + threeCount * 3) break
        fiveCount -= 1
        threeCount = Math.floor((weight-(fiveCount*5))/3)

        if(fiveCount < 0) {
            console.log(-1)
            return
        }
    }

    console.log(fiveCount + threeCount)
})