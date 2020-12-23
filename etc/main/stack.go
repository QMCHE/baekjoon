package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	array []int
}

func (s *stack) push(v int) {
	s.array = append(s.array, v)
}

func (s *stack) pop() int {
	if len(s.array) == 0 {
		return -1
	}

	temp := s.array[len(s.array)-1:][0]
	s.array = s.array[:len(s.array)-1]
	return temp
}

func (s stack) size() int {
	return len(s.array)
}

func (s stack) isEmpty() int {
	if len(s.array) == 0 {
		return 1
	}

	return 0
}

func (s stack) top() int {
	if len(s.array) == 0 {
		return -1
	}

	return s.array[len(s.array)-1]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	array := make([]int, 0)
	stack := stack{array}

	for i := 0; i < n; i++ {
		var command string
		fmt.Fscan(r, &command)

		if command == "push" {
			var value int
			fmt.Fscan(r, &value)
			stack.push(value)
		} else if command == "pop" {
			fmt.Fprint(w, stack.pop(), "\n")
		} else if command == "size" {
			fmt.Fprint(w, stack.size(), "\n")
		} else if command == "empty" {
			fmt.Fprint(w, stack.isEmpty(), "\n")
		} else if command == "top" {
			fmt.Fprint(w, stack.top(), "\n")
		}
	}
}
