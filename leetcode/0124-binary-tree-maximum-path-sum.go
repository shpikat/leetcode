package leetcode

import "math"

func maxPathSum(root *TreeNode) int {
	_, max := findMaxPathSum(root)
	return max
}

func findMaxPathSum(root *TreeNode) (sum int, max int) {
	if root == nil {
		return 0, math.MinInt
	}

	left, leftMax := findMaxPathSum(root.Left)
	right, rightMax := findMaxPathSum(root.Right)

	sum = root.Val
	if left > 0 && left > right {
		sum += left
	} else if right > 0 {
		sum += right
	}

	max = root.Val
	if left > 0 {
		max += left
	}
	if right > 0 {
		max += right
	}
	if leftMax > max {
		max = leftMax
	}
	if rightMax > max {
		max = rightMax
	}

	return
}
