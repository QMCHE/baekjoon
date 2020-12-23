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

	var n, m int
	fmt.Fscan(r, &n, &m)

	for i := 1; i <= n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fprint(w, i)
			for k := 1; k <= n; k++ {
				if i != k {
					fmt.Fprint(w, k)
				}
			}
			fmt.Fprintln(w)
		}
	}
}