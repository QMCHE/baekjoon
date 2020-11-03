const stdin = process.openStdin()

stdin.addListener("data", (d: Buffer) => {
    const n = Number(d.toString())
    console.log(f(1, n))
})

const f = (previousValue: number, currentCount: number): number => {
    if(currentCount === 0) {
        return previousValue
    }

    return f(previousValue*currentCount, currentCount-1)
}