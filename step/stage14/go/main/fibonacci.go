package main

import (
	"bufio"
	"fmt"
	"os"
)

var zero = 0
var one = 0

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var t int
	fmt.Fscan(r, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(r, &n)
		fib(n)
		fmt.Fprintf(w, "%d %d\n", zero, one)
		zero = 0
		one = 0
	}
}

func fib(n int) {
	if n == 0 {
		zero += 1
		return
	} else if n == 1 {
		one += 1
		return
	}
	fib(n - 1)
	fib(n - 2)
}
