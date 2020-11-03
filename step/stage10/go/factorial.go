package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(f(1, n))
}

func f(previousValue, currentCount int) int {
	if currentCount == 0 {
		return previousValue
	}

	return f(previousValue*currentCount, currentCount-1)
}
