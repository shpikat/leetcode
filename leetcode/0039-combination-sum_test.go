package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	testCases := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{
			[]int{2, 3, 6, 7},
			7,
			[][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			[]int{2, 3, 5},
			8,
			[][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			[]int{2},
			1,
			[][]int{},
		},
		{
			[]int{7, 3, 2},
			18,
			[][]int{
				{7, 7, 2, 2},
				{7, 3, 3, 3, 2},
				{7, 3, 2, 2, 2, 2},
				{3, 3, 3, 3, 3, 3},
				{3, 3, 3, 3, 2, 2, 2},
				{3, 3, 2, 2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2, 2, 2, 2, 2},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v (%v)", tc.candidates, tc.target), func(t *testing.T) {
			got := combinationSum(tc.candidates, tc.target)
			if len(got) != len(tc.want) {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			} else {
				want := make(map[string]bool, len(tc.want))
				for _, w := range tc.want {
					sort.Ints(w)
					want[fmt.Sprint(w)] = true
				}

				for _, g := range got {
					sort.Ints(g)
					s := fmt.Sprint(g)
					_, exists := want[s]
					if exists {
						delete(want, s)
					} else {
						t.Errorf("Want: %v. Got: %v", tc.want, got)
						return
					}
				}

				if len(want) != 0 {
					t.Errorf("Want: %v. Got: %v", tc.want, got)
				}
			}
		})
	}
}
