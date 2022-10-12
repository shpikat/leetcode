package leetcode

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)

	top := 0
	right := n - 1
	bottom := m - 1
	left := 0

	for {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		if bottom < top {
			break
		}
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return result
}
