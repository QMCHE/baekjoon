package main

import "fmt"

func main() {
	// n = length
	// x = max number
	var n, x int
	fmt.Scan(&n, &x)

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if x > a {
			fmt.Print(a, " ")
		}
	}
}