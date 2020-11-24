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

	for i := 0; i < n; i++ {
		str := ""
		for j := 0; j < n-i-1; j++ {
			str += " "
		}
		str += "*"
		for j := 0; j < i*2-1; j++ {
			str += " "
		}
		if i != 0 {
			str += "*"
		}
		fmt.Fprint(w, str, "\n")
	}
}