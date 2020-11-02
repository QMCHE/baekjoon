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

	var t int
	fmt.Fscan(r, &t)

	for i := 1; i <= t; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		fmt.Fprintf(w, "Case #%d: %d\n", i, a + b)
	}
}