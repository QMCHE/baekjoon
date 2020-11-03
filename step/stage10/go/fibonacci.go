package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(fib(n))
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-2) + fib(n-1)
}
