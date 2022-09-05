package leetcode

import (
	"fmt"
	"strings"
)

const (
	left  = "({["
	right = ")}]"
)

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	last := len(*s) - 1
	v := (*s)[last]
	*s = (*s)[:last]
	return v
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func isValid(s string) bool {
	stack := Stack{}
	for _, r := range s {
		ch := byte(r)
		index := strings.IndexByte(left, ch)
		if index < 0 {
			index = strings.IndexByte(right, ch)
			if index < 0 {
				// Guaranteed by the constraints not to happen, but we'd better be sure
				fmt.Printf("Unexpected character: %c\n", r)
				return false
			} else if stack.IsEmpty() || stack.Pop() != index {
				return false
			}
		} else {
			stack.Push(index)
		}
	}
	return stack.IsEmpty()
}
