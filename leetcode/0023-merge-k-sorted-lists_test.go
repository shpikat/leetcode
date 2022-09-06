package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	testCases := []struct {
		lists [][]int
		want  []int
	}{
		{
			[][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			[][]int{},
			[]int{},
		},
		{
			[][]int{{}},
			[]int{},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v->%v", tc.lists, tc.want), func(t *testing.T) {
			lists := make([]*ListNode, len(tc.lists))
			for i := range tc.lists {
				lists[i] = toList(tc.lists[i])
			}
			got := fromList(mergeKLists(lists))
			if len(got) != len(tc.want) {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			} else {
				for i := range got {
					if got[i] != tc.want[i] {
						t.Errorf("Want: %v. Got: %v", tc.want, got)
						break
					}
				}
			}
		})
	}
}
