package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var ground [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var a int
			fmt.Fscan(r, &a)
			ground[i][j] = a
		}
	}

	for {
		for i, arr := range ground {
			for j, num := range arr {
				if num == 0 {
					x, xv := remainNumInX(ground, i)
					if x {
						ground[i][j] = xv
					} else {
						y, yv := remainNumInY(ground, j)
						if y {
							ground[i][j] = yv
						} else {
							b, bv := remainNumInBox(ground, i/3, j/3)
							if b {
								ground[i][j] = bv
							}
						}
					}
				}
			}
		}
		if isComplete(ground) {
			break
		}
	}

	for _, arr := range ground {
		for _, num := range arr {
			fmt.Fprintf(w, "%d ", num)
		}
		fmt.Fprint(w, "\n")
	}

	return
}

func remainNumInX(arr [9][9]int, index int) (bool, int) {
	array := [9]int{0}

	for i := 0; i < 9; i++ {
		if arr[index][i] != 0 {
			array[arr[index][i]-1]++
		}
	}

	remain := remainNumber(array, func(arr [9]int, num int) bool {
		return num != 1
	})

	return len(remain) == 1, remain[0]
}

func remainNumInY(arr [9][9]int, index int) (bool, int) {
	array := [9]int{0}

	for i := 0; i < 9; i++ {
		if arr[i][index] != 0 {
			array[arr[i][index]-1]++
		}
	}

	remain := remainNumber(array, func(arr [9]int, num int) bool {
		return num != 1
	})

	return len(remain) == 1, remain[0]
}

func remainNumInBox(arr [9][9]int, x, y int) (bool, int) {
	array := [9]int{0}

	x *= 3
	y *= 3

	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if arr[i][j] != 0 {
				array[arr[i][j]-1]++
			}
		}
	}

	remain := remainNumber(array, func(arr [9]int, num int) bool {
		return num != 1
	})

	return len(remain) == 1, remain[0]
}

func isComplete(arr [9][9]int) bool {
	for _, a := range arr {
		for _, num := range a {
			if num == 0 {
				return false
			}
		}
	}

	return true
}

func remainNumber(arr [9]int, test func([9]int, int) bool) []int {
	newArr := make([]int, 0)

	for i, num := range arr {
		if test(arr, num) {
			newArr = append(newArr, i+1)
		}
	}

	return newArr
}