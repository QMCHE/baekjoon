const stdin = process.openStdin()

stdin.addListener("data", (d: Buffer) => {
    const str = d.toString()
    console.log(str)
})