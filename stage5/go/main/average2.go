package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	var sum float64 = 0
	var max float64 = 0

	for i := 0; i < n; i++ {
		var a float64
		fmt.Fscan(r, &a)
		sum += a
		if max < a {
			max = a
		}
	}

	fmt.Fprintf(w, "%.6f", ((sum/max)*100)/float64(n))
}