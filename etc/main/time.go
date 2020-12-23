package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	//a, b, c := 0, 0, 0

	startA := time.Now()
	for i := 0; i < 10; i++ {
		fmt.Print("1")
	}
	fmt.Print("\n\nA: ", time.Since(startA), "\n\n")

	time.Sleep(time.Millisecond)

	startB := time.Now()
	//r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < 10; i++ {
		fmt.Fprint(w, "1")
	}
	fmt.Print("\n\nB: ", time.Since(startB), "\n\n")

	//time.Sleep(time.Millisecond)

	//startC := time.Now()
	//for i := 0; i < 100000; i++ {
	//	c++
	//}
	//fmt.Print("C: ", time.Since(startC), "\n")
}