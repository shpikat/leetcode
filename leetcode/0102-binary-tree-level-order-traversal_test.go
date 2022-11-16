package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	testCases := []struct {
		tree string
		want [][]int
	}{

		{
			"[3,9,20,null,null,15,7]",
			[][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			"[1]",
			[][]int{{1}},
		},
		{
			"[]",
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.tree), func(t *testing.T) {
			got := levelOrder(toBinaryTree(tc.tree))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			}
		})
	}
}
