package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	var num string
	fmt.Fscan(r, &n, &num)
	sum := 0

	array := strings.Split(num, "")

	for _, number := range array {
		a, _ := strconv.Atoi(number)
		sum += a
	}

	fmt.Fprint(w, sum)
}