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

	for i := n; i > 0; i-- {
		str := ""
		for j := 0; j < i; j++ {
			str += "*"
		}
		fmt.Fprint(w, str+"\n")
	}
}