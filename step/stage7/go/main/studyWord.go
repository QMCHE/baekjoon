package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var str string
	fmt.Fscan(r, &str)
	str = strings.ToUpper(str)
	strArr := strings.Split(str, "")

	var count = [26]int{0}

	max := 0
	answer := ""

	for _, char := range strArr {
		index := char[0] - 'A'
		count[index]++

		if count[index] > max {
			max = count[index]
			answer = char
		} else if count[index] == max {
			answer = "?"
		}
	}

	fmt.Fprint(w, answer)
}