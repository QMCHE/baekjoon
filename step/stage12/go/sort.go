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

	var n int
	fmt.Fscan(r, &n)
	var array []int

	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(r, &a)
		array = append(array, a)
	}

	for i := 1; i < len(array); i++ {
		temp := array[i]
		array = append(array[:i], array[i+1:]...)
		for j := i - 1; j >= 0; j-- {
			if temp > array[j] || j == 0 {
				fmt.Fprint(w, array[:j])
				fmt.Fprint(w, array[j:])
				a := append(array[:j], temp)
				fmt.Fprint(w, a)
				array = append(a, array[j:]...)
				fmt.Fprintln(w, array)
				break
			}
		}
	}

	fmt.Fprint(w, array)
}
