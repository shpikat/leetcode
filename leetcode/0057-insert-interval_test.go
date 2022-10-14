package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{
			[][]int{{1, 3}, {6, 9}},
			[]int{2, 5},
			[][]int{{1, 5}, {6, 9}},
		},
		{
			[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			[]int{4, 8},
			[][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			[][]int{{2, 3}},
			[]int{1, 2},
			[][]int{{1, 3}},
		},
		{
			[][]int{{2, 3}},
			[]int{3, 4},
			[][]int{{2, 4}},
		},
		{
			[][]int{{2, 3}},
			[]int{1, 4},
			[][]int{{1, 4}},
		},
		{
			[][]int{{1, 4}},
			[]int{2, 3},
			[][]int{{1, 4}},
		},
		{
			[][]int{{3, 4}},
			[]int{1, 2},
			[][]int{{1, 2}, {3, 4}},
		},
		{
			[][]int{{1, 2}},
			[]int{3, 4},
			[][]int{{1, 2}, {3, 4}},
		},
		{
			[][]int{{1, 2}, {5, 6}},
			[]int{3, 4},
			[][]int{{1, 2}, {3, 4}, {5, 6}},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v-%v", tc.intervals, tc.newInterval), func(t *testing.T) {
			got := insert(tc.intervals, tc.newInterval)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			}
		})
	}
}
