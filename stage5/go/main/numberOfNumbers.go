package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, b, c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	sum := strconv.Itoa(a * b * c)
	fmt.Println(sum)

	for i := 0; i < 10; i++ {
		fmt.Println(strings.Count(sum, strconv.Itoa(i)))
	}
}