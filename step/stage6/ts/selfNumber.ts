const d = (a: number): number => {
    if(a < 10) {
        return a * 2
    }

    const num = a.toString()
    for(let i=0; i<num.length; i++) {
        a += Number(num.charAt(i))
    }

    return a
}

let array = new Array()

for(let i=0; i<=9972; i++) {
    const newNum = d(i)

    array[newNum] = newNum
}

for(let i=0; i<array.length; i++) {
    if(i > 10000) {
        break
    }
    array[i] === undefined && console.log(i)
}