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
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		fmt.Fprintln(w, a+b)
	}
}