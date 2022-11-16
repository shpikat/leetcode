package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	testCases := []struct {
		tree string
		want int
	}{
		{
			"[1,2,3]",
			6,
		},
		{
			"[-10,9,20,null,null,15,7]",
			42,
		},
		{
			"[5,4,8,11,null,13,4,7,2,null,null,null,1]",
			48,
		},
		{
			"[-3]",
			-3,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.tree), func(t *testing.T) {
			got := maxPathSum(toBinaryTree(tc.tree))
			if got != tc.want {
				t.Errorf("Want: %d. Got: %d", tc.want, got)
			}
		})
	}
}
