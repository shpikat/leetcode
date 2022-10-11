package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		strs []string
		want [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			[]string{""},
			[][]string{{""}},
		},
		{
			[]string{"a"},
			[][]string{{"a"}},
		},
		{
			[]string{"aab", "abb", "aba", "baa"},
			[][]string{{"aab", "baa", "aba"}, {"abb"}},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprint(tc.strs), func(t *testing.T) {
			got := groupAnagrams(tc.strs)

			strs := make(map[string]Void, len(tc.strs))
			for _, str := range tc.strs {
				strs[str] = void
			}

			for _, group := range got {
				expected := map[string]Void{}
				for _, s := range group {
					_, exists := strs[s]
					if !exists {
						t.Errorf("Want: %v. Got: %v", tc.want, got)
						return
					}
					// Use a different sorting approach
					bs := []byte(s)
					sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
					expected[string(bs)] = void
				}
				if len(expected) != 1 {
					t.Errorf("Want: %v. Got: %v", tc.want, got)
					return
				}
			}
		})
	}
}
