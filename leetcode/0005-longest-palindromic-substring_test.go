package leetcode

import (
	"strings"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		s    string
		want string
	}{
		{"", ""},
		{"a", "a"},
		{"bb", "bb"},
		{"ccc", "ccc"},
		{"babad", "aba"},
		{"cbbd", "bb"},
		{"1caroracat2wasitacaroracatisaw3", "wasitacaroracatisaw"},
		{"abcbabcba", "abcbabcba"},
		{"aabbbccc", "bbb"},
		{"aaccccbbbbcccc", "ccccbbbbcccc"},
		{"abbcccbbbcaaccbababcbcabca", "bbcccbb"},
	}
	for _, tc := range testCases {
		t.Run(tc.s, func(t *testing.T) {
			got := longestPalindrome(tc.s)
			if !strings.Contains(tc.s, got) {
				t.Errorf("Not a substring: %s <-> %s", tc.s, got)
			}
			for i, j := 0, len(got)-1; i < len(got)/2; i, j = i+1, j-1 {
				if got[i] != got[j] {
					t.Errorf("Not palindromic: %s-%s-%s", got[:i], got[i:j+1], got[j+1:])
					break
				}
			}
			if len(got) < len(tc.want) {
				t.Errorf("Not the longest: %s > %s", tc.want, got)
			}
		})
	}
}
