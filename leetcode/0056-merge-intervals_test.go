package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		intervals [][]int
		want      [][]int
	}{
		{
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
		{
			[][]int{{1, 6}, {2, 3}, {2, 4}, {5, 7}, {8, 10}},
			[][]int{{1, 7}, {8, 10}},
		},
		{
			[][]int{{1, 4}, {0, 4}},
			[][]int{{0, 4}},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.intervals), func(t *testing.T) {
			got := merge(tc.intervals)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			}
		})
	}
}
