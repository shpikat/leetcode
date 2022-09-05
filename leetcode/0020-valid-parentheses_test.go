package leetcode

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"(((([][]{{{[{[{}][]}]}}}[])))())", true},
		{"(((([][]{{{[{[{}][]}]}}}[])))()", false},
		{"((([][]{{{[{[{}][]}]}}}[])))())", false},
	}
	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			got := isValid(tc.s)
			if got != tc.want {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			}
		})
	}
}
