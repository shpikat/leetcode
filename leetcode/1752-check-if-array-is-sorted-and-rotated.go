package leetcode

func check(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			if nums[0] < nums[len(nums)-1] {
				return false
			}
			for j := i + 1; j < len(nums)-1; j++ {
				if nums[j+1] < nums[j] {
					return false
				}
			}
		}
	}
	return true
}
