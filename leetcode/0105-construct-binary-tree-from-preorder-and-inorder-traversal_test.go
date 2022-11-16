package leetcode

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	testCases := []struct {
		preorder []int
		inorder  []int
		want     string
	}{
		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			"[3,9,20,null,null,15,7]",
		},
		{
			[]int{-1},
			[]int{-1},
			"[-1]",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v-%v", tc.preorder, tc.inorder), func(t *testing.T) {
			got := fromBinaryTree(buildTree(tc.preorder, tc.inorder))
			if got != tc.want {
				t.Errorf("Want: %s. Got: %s", tc.want, got)
			}
		})
	}
}
