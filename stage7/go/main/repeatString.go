package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rr := bufio.NewReader(os.Stdin)
	ww := bufio.NewWriter(os.Stdout)
	defer ww.Flush()

	var t int
	fmt.Fscan(rr, &t)

	for i := 0; i < t; i++ {
		var r int
		var str string
		fmt.Fscan(rr, &r, &str)

		strArray := strings.Split(str, "")

		for _, char := range strArray {
			for j := 0; j < r; j++ {
				fmt.Fprint(ww, char)
			}
		}
		fmt.Fprintln(ww)
	}
}