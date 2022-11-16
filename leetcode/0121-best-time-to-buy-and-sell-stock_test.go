package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			[]int{7, 1, 5, 0, 3},
			4,
		},
		{
			[]int{7, 1, 5, 0, 6},
			6,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.prices), func(t *testing.T) {
			got := maxProfit(tc.prices)
			if got != tc.want {
				t.Errorf("Want: %d. Got: %d", tc.want, got)
			}
		})
	}
}
