package leetcode

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	const maxLength = 17
	for length := 1; length <= maxLength; length++ {
		nums := make([]int, length)
		for pivotIndex := 1; pivotIndex < length; pivotIndex++ {
			for i := range nums {
				nums[i] = (i + pivotIndex) % length
			}
			for want := 0; want < length; want++ {
				target := nums[want]
				t.Run(fmt.Sprintf("%v (%d)", nums, target), func(t *testing.T) {
					got := search(nums, target)
					if got != want {
						t.Errorf("Want: %v. Got: %v", want, got)
					}
				})
			}
		}
	}
}
