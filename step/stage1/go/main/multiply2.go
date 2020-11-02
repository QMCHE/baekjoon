package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	sum := a * b
	units := a * (b%10)
	b /= 10
	tens := a * (b%10)
	b /= 10
	hundreds := a * (b%10)

	fmt.Println(units)
	fmt.Println(tens)
	fmt.Println(hundreds)
	fmt.Print(sum)
}