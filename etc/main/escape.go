package main

import (
	"fmt"
)

func main() {
	var x, y, w, h int
	fmt.Scan(&x, &y, &w, &h)

	min := 1000

	if min > h - y {
		min = h - y
	}
	if min > y {
		min = y
	}
	if min > x {
		min = x
	}
	if min > w - x {
		min = w - x
	}
	fmt.Print(min)
}