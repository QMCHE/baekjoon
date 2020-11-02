package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var array = [10000]int{0}
	for i := 1; i <= 10000; i++ {
		newNum := d(i)

		if newNum >= 10000 {
			break
		}

		array[newNum] = newNum
	}

	for _, num := range array {
		if num != 0 {
			fmt.Fprintln(w, num)
		}
	}
}

func d(n int) int {
	sum := n

	for {
		sum += n%10
		n /= 10

		if n <= 0 {
			break
		}
	}

	return sum
}