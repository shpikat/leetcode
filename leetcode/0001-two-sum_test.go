package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.nums), func(t *testing.T) {
			got := twoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			}
		})
	}
}
