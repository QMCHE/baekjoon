package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	num := n
	count := 0

	for {
		count++
		units := num%10
		hundreds := num/10

		num = (units * 10) + ((units + hundreds)%10)

		if n == num {
			break
		}
	}

	fmt.Print(count)
}