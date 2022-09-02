package leetcode

func maxArea(height []int) int {
	max := 0
	left, right := 0, len(height)-1
	for left < right {
		var h int
		width := right - left
		if height[left] <= height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		current := h * width
		if current > max {
			max = current
		}
	}

	return max
}
