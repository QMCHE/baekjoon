const stdin = process.openStdin()

stdin.addListener("data", (d: Buffer) => {
    let str = d.toString()
    
    for(let i=1; i<=Math.ceil(str.length/10); i++) {
        console.log(str.substring((i-1)*10, i*10))
    }
})