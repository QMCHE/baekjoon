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

	var s string
	fmt.Fscan(r, &s)

	for i := 0; i < 26; i++ {
		fmt.Fprint(w, strings.IndexRune(s, rune(97+i)))
		fmt.Fprint(w, " ")
	}
}