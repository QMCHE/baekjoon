let a = []

for (let i = 0; i < 100000; i++) {
    let rand = Math.floor(Math.random() * 9999 + 1).toString().split("")
    for (let j = 0; j < rand.length; j++) {
        a[Number(rand[j])] = ((a[Number(rand[j])] || 0) + 1)
    }
}

console.log(a)