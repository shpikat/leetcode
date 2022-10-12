package leetcode

func maxSubArray(nums []int) int {
	sum := 0
	max := nums[0]
	for _, n := range nums {
		sum += n
		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return max
}
