package leetcode

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	testCases := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCCED",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"SEE",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCB",
			false,
		},
		{
			[][]byte{{'a'}},
			"a",
			true,
		},
		{
			[][]byte{{'a'}, {'a'}},
			"A",
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s-%s", tc.board, tc.word), func(t *testing.T) {
			got := exist(tc.board, tc.word)
			if got != tc.want {
				t.Errorf("Want: %t. Got: %t", tc.want, got)
			}
		})
	}
}
