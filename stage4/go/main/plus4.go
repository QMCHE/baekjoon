package main

import (
	"fmt"
)

func main() {
	for {
		var a, b int
		_, err := fmt.Scan(&a, &b)

		if err != nil {
			break
		}

		fmt.Println(a+b)
	}
}