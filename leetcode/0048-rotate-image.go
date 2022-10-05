package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)
	last := n - 1
	for start := 0; start < n/2; start++ {
		end := last - start
		for i := start; i < end; i++ {
			k := last - i
			matrix[i][end], matrix[end][k], matrix[k][start], matrix[start][i] = matrix[start][i], matrix[i][end], matrix[end][k], matrix[k][start]
		}
	}
}
