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
	fmt.Fscan(r, &n)

	array := make([]int, n)

	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(r, &a)
		array[i] = a
	}

	fmt.Fscan(r, &m)

	for i := 0; i < m; i++ {
		var a int
		isExist := 0
		fmt.Fscan(r, &a)

		for _, num := range array {
			if a == num {
				isExist = 1
				break
			}
		}

		fmt.Fprint(w, isExist, "\n")
	}
}
