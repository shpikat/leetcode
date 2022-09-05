package leetcode

func twoSum(nums []int, target int) []int {
	known := make(map[int]int, len(nums))
	for i, n := range nums {
		j, exists := known[target-n]
		if exists {
			return []int{i, j}
		} else {
			known[n] = i
		}
	}
	return []int{}
}
