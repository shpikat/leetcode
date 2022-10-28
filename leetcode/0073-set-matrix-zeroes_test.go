package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			[][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
		{
			[][]int{{1, 0, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]int{{0, 0, 0}, {1, 0, 1}, {1, 0, 1}},
		},
		{
			[][]int{{1, 0, 1}, {1, 1, 1}, {1, 1, 1}},
			[][]int{{0, 0, 0}, {1, 0, 1}, {1, 0, 1}},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.matrix), func(t *testing.T) {
			setZeroes(tc.matrix)
			if !reflect.DeepEqual(tc.matrix, tc.want) {
				t.Errorf("Want: %v. Got: %v", tc.want, tc.matrix)
			}
		})
	}
}
