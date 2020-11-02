package main

import "fmt"

func main() {
	var h, m int
	fmt.Scan(&h, &m)

	m -= 45

	if m < 0 {
		h -= 1
		m += 60
	}

	if h < 0 {
		h += 24
	}

	fmt.Print(h, m)
}