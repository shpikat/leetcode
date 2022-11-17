package leetcode

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		s    string
		want bool
	}{
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"race a car",
			false,
		},
		{
			" ",
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.s), func(t *testing.T) {
			got := isPalindrome(tc.s)
			if got != tc.want {
				t.Errorf("Want: %t. Got: %t", tc.want, got)
			}
		})
	}
}
