package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	testCases := []struct {
		matrix [][]int
		want   []int
	}{
		{[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			[]int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8},
		},
		{
			[][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
			},
			[]int{1, 2, 3, 4, 5, 10, 15, 14, 13, 12, 11, 6, 7, 8, 9},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.matrix), func(t *testing.T) {
			got := spiralOrder(tc.matrix)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Want: %v. Got: %v.", tc.want, got)
			}
		})
	}
}
