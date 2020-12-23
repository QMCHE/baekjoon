package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n < 2 {
		fmt.Print(n)
		return
	}

	previous, current := 0, 1

	for i := 1; i < n; i++ {
		temp := current
		current += previous
		previous = temp
	}

	fmt.Print(current)
}