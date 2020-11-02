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

	str, _ := r.ReadString('\n')
	fmt.Fprint(w, len(strings.Split(str, " ")))
}