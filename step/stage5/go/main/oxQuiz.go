package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		var str string
		fmt.Fscan(r, &str)
		results := strings.Split(str, "")
		sum, score := 0, 0

		for _, result := range results {
			if result == "O" {
				score += 1
				sum += score
			} else {
				score = 0
			}
		}

		fmt.Fprintln(w, sum)
	}
}