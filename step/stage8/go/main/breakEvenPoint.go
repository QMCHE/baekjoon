package main

import "fmt"

import "bufio"

import "os"

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b, c int
	fmt.Fscan(r, &a, &b, &c)

	// 노트북 1대 당 순수익 = 노트북 가격 - 가변 비용
	netProfit := c - b

	if netProfit <= 0 {
		// 얻는 수익이 없을 경우 손익분기점이 발생하지 않는다.
		fmt.Fprint(w, -1)
	} else {
		// 얻는 수익으로 고정 비용을 충당한 경우 손익분기점이 발생한다.
		fmt.Fprint(w, (a/netProfit)+1)
	}
}
