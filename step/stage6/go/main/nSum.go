package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	intArray := []int{1, 2, 3, 4, 5, 6}
	fmt.Fprint(w, sum(intArray))
}

func sum(a []int) int {
	sum := 0

	for _, num := range a {
		sum += num
	}

	return sum
}