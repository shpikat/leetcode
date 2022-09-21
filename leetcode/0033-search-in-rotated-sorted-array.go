package leetcode

func search(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		k := (i + j) / 2
		if target == nums[k] {
			return k
		}

		isLeftPartMonotonic := nums[k] > nums[i]
		isTargetOnTheRightSide := target < nums[i]
		// Can be split back into nested conditions, which will improve readability a little.
		if (target > nums[k] && (isLeftPartMonotonic || isTargetOnTheRightSide)) ||
			(isLeftPartMonotonic && isTargetOnTheRightSide) {
			i = k + 1
		} else {
			j = k
		}
	}
	return -1
}
