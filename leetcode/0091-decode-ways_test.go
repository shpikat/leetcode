package leetcode

import (
	"fmt"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	testCases := []struct {
		s    string
		want int
	}{
		{
			"12",
			2,
		},
		{
			"226",
			3,
		},
		{
			"06",
			0,
		},
		{
			"10",
			1,
		},
		{
			"2611055971756562",
			4,
		},
		{
			"111111111111111111111111111111111111111111111",
			1836311903,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.s), func(t *testing.T) {
			got := numDecodings(tc.s)
			if got != tc.want {
				t.Errorf("Want: %d. Got: %d", tc.want, got)
			}
		})
	}
}
