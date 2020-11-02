package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r:= bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	max := 0
	index := 0
	array := make([]int, 9)

	for i := 0; i < 9; i++ {
		var num int
		fmt.Fscan(r, &num)
		array[i] = num

		if max < num {
			max = num
			index = i+1
		}
	}

	fmt.Fprintln(w, max)
	fmt.Fprintln(w, index)
}