package leetcode

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {
	testCases := []struct {
		s    string
		t    string
		want string
	}{
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"a",
			"a",
			"a",
		},
		{
			"a",
			"aa",
			"",
		},
		{
			"a",
			"A",
			"",
		},
		{
			"ab",
			"a",
			"a",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s-%s", tc.s, tc.t), func(t *testing.T) {
			got := minWindow(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("Want: %s. Got: %s", tc.want, got)
			}
		})
	}
}
