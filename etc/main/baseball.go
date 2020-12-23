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

	for i := 0; i < t; i++ {
		yonsei, korea := 0, 0

		for j := 0; j < 9; j++ {
			var y, k int
			fmt.Fscan(r, &y, &k)
			if y > k {
				yonsei++
			} else if y < k {
				korea++
			}
		}

		if yonsei > korea {
			fmt.Fprint(w, "Yonsei\n")
		} else if yonsei < korea {
			fmt.Fprint(w, "Korea\n")
		} else {
			fmt.Fprint(w, "Draw\n")
		}
	}
}