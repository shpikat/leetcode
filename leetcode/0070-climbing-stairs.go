package leetcode

func climbStairs(n int) int {
	previous, current := 0, 1

	for step := 1; step <= n; step++ {
		previous, current = current, previous+current

	}

	return current
}
