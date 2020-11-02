package main

import (
	"fmt"
)

func main() {
	var c int
	fmt.Scan(&c)

	for i := 0; i < c; i++ {
		var n int
		fmt.Scan(&n)

		array := make([]int, n)

		for j, _ := range array {
			var score int
			fmt.Scan(&score)
			array[j] = score
		}

		avg := sum(array)/n
		count := 0

		for j, _ := range array {
			if array[j] > avg {
				count++
			}
		}

		percent := float64(count) / float64(n)
		fmt.Printf("%.3f%s\n", percent*100, "%")
	}
}

func sum(arr []int) int {
	a := 0

	for i, _ := range arr {
		a += arr[i]
	}

	return a
}