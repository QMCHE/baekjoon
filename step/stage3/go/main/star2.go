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

	var str string
	for i := 0; i < n; i++ {
		str += "*"
		for j := n-1; j > i; j-- {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, str + "\n")
	}
}