package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var mon, day int
	fmt.Scan(&mon, &day)
	t := time.Date(2007, time.Month(mon), day, 0, 0, 0, 0, time.UTC)
	fmt.Print(strings.ToUpper(t.Weekday().String()[:3]))
}