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
		fmt.Fprint(w, str + "\n")
	}
}