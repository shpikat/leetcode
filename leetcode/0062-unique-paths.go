package leetcode

func uniquePaths(m int, n int) int {
	// m is chosen to be the greater number of the two
	if n > m {
		m, n = n, m
	}

	// The number of steps is one short of the number of cells
	m--

	// Optimized calculation of the number of permutations: (m+n)! / m!*n!
	x := 1
	for i := 1; i < n; i++ {
		x = (x * (m + i)) / i
	}
	return x
}
