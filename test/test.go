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

	for {
		var str string

		if err != nil {
			break
		}
		fmt.Fprint(w, str+"\n")
	}
}
