package leetcode

import (
	"strconv"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		n    int
		want int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{1, 1},
	}

	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.n), func(t *testing.T) {
			got := climbStairs(tc.n)
			if got != tc.want {
				t.Errorf("Want: %d. Got: %d", tc.want, got)
			}
		})
	}
}
