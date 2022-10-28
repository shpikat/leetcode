package leetcode

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	testCases := []struct {
		m    int
		n    int
		want int
	}{
		{3, 7, 28},
		{3, 2, 3},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%dx%d", tc.m, tc.n), func(t *testing.T) {
			got := uniquePaths(tc.m, tc.n)
			if got != tc.want {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			}
		})
	}
}
