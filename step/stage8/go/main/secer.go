package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	five := n / 5
	three := (n - (five * 5)) / 3

	for {
		if n == five*5+three*3 {
			break
		}
		five--
		if five < 0 {
			fmt.Print(-1)
			return
		}
		three = (n - five*5) / 3
	}

	fmt.Print(five + three)
}
