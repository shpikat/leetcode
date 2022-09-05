package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		list1 []int
		list2 []int
		want  []int
	}{
		{
			[]int{1, 2, 4},
			[]int{1, 2, 3},
			[]int{1, 1, 2, 2, 3, 4},
		},
		{
			[]int{},
			[]int{},
			[]int{},
		},
		{
			[]int{},
			[]int{0},
			[]int{0},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v+%v=%v", tc.list1, tc.list2, tc.want), func(t *testing.T) {
			got := fromList(mergeTwoLists(toList(tc.list1), toList(tc.list2)))
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
