package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		tree string
		want int
	}{

		{
			"[3,9,20,null,null,15,7]",
			3,
		},
		{
			"[1,null,2]",
			2,
		},
		{
			"[]",
			0,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.tree), func(t *testing.T) {
			got := maxDepth(toBinaryTree(tc.tree))
			if got != tc.want {
				t.Errorf("Want: %d. Got: %d", tc.want, got)
			}
		})
	}
}
