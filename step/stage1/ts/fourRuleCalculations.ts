const stdin = process.openStdin()

stdin.addListener("data", (d: Buffer) => {
    const [a, b] = d.toString().split(" ").map(Number)
    console.log(a+b)
    console.log(a-b)
    console.log(a*b)
    console.log(Math.floor(a/b))
    console.log(a%b)
})