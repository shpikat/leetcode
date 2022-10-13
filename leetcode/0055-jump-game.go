package leetcode

func canJump(nums []int) bool {
	reaching := 0
	for i := 0; i <= reaching && reaching < len(nums); i++ {
		next := i + nums[i]
		if reaching < next {
			reaching = next
		}
	}
	return reaching >= len(nums)-1
}
