package main

import "fmt"

func main() {
	var year int
	fmt.Scan(&year)

	if year%4 == 0 {
		if year%100 != 0 {
			fmt.Print(1)
		} else if year%400 == 0 {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	} else {
		fmt.Print(0)
	}
}