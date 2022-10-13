package leetcode

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			[]int{2, 3, 1, 1, 4},
			true,
		},
		{[]int{3, 2, 1, 0, 4},
			false,
		},
		{
			[]int{2, 0},
			true,
		},
		{
			[]int{1},
			true,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.nums), func(t *testing.T) {
			got := canJump(tc.nums)
			if got != tc.want {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			}
		})
	}
}
