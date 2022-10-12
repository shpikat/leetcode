package leetcode

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		nums []int
		want int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{5, 4, -1, 7, 8},
			23,
		},
		{
			[]int{-2, -1},
			-1,
		},
		{
			[]int{-2, -3, -1},
			-1,
		},
		{
			[]int{-3, -2, 0, -1},
			0,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.nums), func(t *testing.T) {
			got := maxSubArray(tc.nums)
			if got != tc.want {
				t.Errorf("Want: %d. Got: %d", tc.want, got)
			}
		})
	}
}
