const stdin = process.openStdin()

stdin.addListener("data", (d: Buffer) => {
    const n = Number(d.toString())
    console.log(fib(n))
})

const fib = (n: number): number => {
    if(n < 2) {
        return n
    }

    return fib(n-2) + fib(n-1)
}