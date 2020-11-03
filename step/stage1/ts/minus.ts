const stdin = process.openStdin()

stdin.addListener("data", d => {
    const [a, b] = d.toString().split(" ").map(Number)
    console.log(a-b)
})