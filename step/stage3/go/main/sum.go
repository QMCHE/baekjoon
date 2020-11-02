package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	sum := 0
	for i := 1; i <= a; i++ {
		sum += i
	}

	fmt.Println(sum)
}