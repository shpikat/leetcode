package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		list []int
		n    int
		want []int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{1, 2, 3, 5},
		},
		{
			[]int{1},
			1,
			[]int{},
		},
		{
			[]int{1, 2},
			1,
			[]int{1},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v - %d", tc.list, tc.n), func(t *testing.T) {
			got := fromList(removeNthFromEnd(toList(tc.list), tc.n))
			if len(got) != len(tc.want) {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("Want: %v. Got: %v", tc.want, got)
				}
			}
		})
	}
}
