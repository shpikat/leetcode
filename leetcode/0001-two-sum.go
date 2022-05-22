package leetcode

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)
	left := 0
	right := len(sorted) - 1
	for {
		leftValue := sorted[left]
		rightValue := sorted[right]
		sum := leftValue + rightValue
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			answer := make([]int, 0, 2)
			for i, value := range nums {
				if value == leftValue || value == rightValue {
					answer = append(answer, i)
					if len(answer) == 2 {
						return answer
					}
				}
			}
		}
	}
}
