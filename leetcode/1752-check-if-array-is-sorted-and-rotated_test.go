package leetcode

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	testCases := []struct {
		nums []int
		want bool
	}{
		{[]int{3, 4, 5, 1, 2}, true},
		{[]int{2, 1, 3, 4}, false},
		{[]int{1, 2, 3}, true},
		{[]int{1, 1, 3, 4, 4}, true},
		{[]int{1, 2, 1, 3, 4, 4}, false},
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 1, 2}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.nums), func(t *testing.T) {
			got := check(tc.nums)
			if got != tc.want {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			}
		})
	}
}
