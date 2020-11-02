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

	var array []int

	for i := 0; i < 10; i++ {
		var a int
		fmt.Fscan(r, &a)

		if All(array, func(v int) bool {
			return a%42 != v
		}) {
			array = append(array, a%42)
		}
	}

	fmt.Fprint(w, len(array))
}

func All(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}