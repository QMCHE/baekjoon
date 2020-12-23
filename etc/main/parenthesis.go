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
		var str string
		fmt.Fscan(r, &str)

		open := 0
		isCorrect := true
		for j := 0; j < len(str); j++ {
			if str[j:j+1] == "(" {
				open++
			} else {
				open--
			}
			if open < 0 {
				isCorrect = false
			}
		}

		if isCorrect && open == 0 {
			fmt.Fprint(w, "YES\n")
		} else {
			fmt.Fprint(w, "NO\n")
		}
	}
}