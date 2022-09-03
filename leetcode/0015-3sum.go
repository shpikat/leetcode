package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	triplets := make([][]int, 0, 16)

	sort.Ints(nums)
	index := 0
	end := len(nums) - 1

	for index < end {
		// For each element we try to find the pair that sums it up to zero
		expected := -nums[index]
		// Early exit, as positive numbers won't sum up to zero
		if expected < 0 {
			break
		}
		left, right := index+1, end
		for left < right {
			s := nums[left] + nums[right]
			if s < expected {
				left++
				// Skip possible duplicate triplets
				for nums[left-1] == nums[left] && left < right {
					left++
				}
			} else if s > expected {
				right--
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			} else {
				triplets = append(triplets, []int{nums[index], nums[left], nums[right]})
				left++
				for nums[left-1] == nums[left] && left < right {
					left++
				}
				right--
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			}
		}
		index++
		for nums[index-1] == nums[index] && index < end {
			index++
		}
	}
	return triplets
}
