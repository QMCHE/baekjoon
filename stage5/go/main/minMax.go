package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(r, &n)

	min := 1000000
	max := -1000000

	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(r, &a)
		if min > a {
			min = a
		}
		if max < a {
			max = a
		}
	}

	fmt.Fprint(w, min, max)
	w.Flush()
}