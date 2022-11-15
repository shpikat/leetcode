package leetcode

import (
	"fmt"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	testCases := []struct {
		p    string
		q    string
		want bool
	}{

		{
			"[1,2,3]",
			"[1,2,3]",
			true,
		},
		{
			"[1,2]",
			"[1,null,2]",
			false,
		},
		{
			"[1,2,1]",
			"[1,1,2]",
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s-%s", tc.p, tc.q), func(t *testing.T) {
			got := isSameTree(toBinaryTree(tc.p), toBinaryTree(tc.q))
			if got != tc.want {
				t.Errorf("Want: %t. Got: %t", tc.want, got)
			}
		})
	}
}
