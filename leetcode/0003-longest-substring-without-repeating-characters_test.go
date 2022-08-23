package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		s    string
		want int
	}{
		{"", 0},
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"abcdefa", 6},
		{"abcdefb", 6},
	}
	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.s)
			if got != tc.want {
				t.Errorf("Want: %v. Got: %v", tc.want, got)
			}
		})
	}
}
