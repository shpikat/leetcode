package leetcode

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{[]int{0, 0}, [][]int{}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{
			{0, 0, 0},
		}},
		{[]int{-1, 0, 1, 1}, [][]int{
			{-1, 0, 1},
		}},
		{[]int{-1, -1, 0, 0, 1, 1}, [][]int{
			{-1, 0, 1},
		}},
		{[]int{1, 1, -2}, [][]int{
			{-2, 1, 1},
		}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.nums), func(t *testing.T) {
			got := threeSum(tc.nums)
			for i := range got {
				sum := 0
				for _, x := range got[i] {
					sum += x
				}
				if sum != 0 {
					t.Errorf("Triplet does not sum up to zero: %v", got[i])
				}
			}

			if len(got) != len(tc.want) {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			}
		})
	}
}
