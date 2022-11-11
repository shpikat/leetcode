package leetcode

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	testCases := []struct {
		tree string
		want bool
	}{

		{
			"[2, 1, 3]",
			true,
		},
		{
			"[5, 1, 4, null, null, 3, 6]",
			false,
		}}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.tree), func(t *testing.T) {
			got := isValidBST(toBinaryTree(tc.tree))
			if got != tc.want {
				t.Errorf("Want: %t. Got: %t", tc.want, got)
			}
		})
	}
}
