package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 8, 7, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 6, 7, 2, 5, 4, 8, 3, 7}, 42},
		{[]int{1, 1}, 1},
		{[]int{0, 0}, 0},
		{[]int{0, 1, 0}, 0},
		{[]int{0, 1, 2, 0}, 1},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.height), func(t *testing.T) {
			got := maxArea(tc.height)
			if got != tc.want {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			}
		})
	}
}
