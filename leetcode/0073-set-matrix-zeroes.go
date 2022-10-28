package leetcode

func setZeroes(matrix [][]int) {
	isFirstRowZero := false
	isFirstColumnZero := false

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					isFirstRowZero = true
				}
				if j == 0 {
					isFirstColumnZero = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	if isFirstRowZero {
		for i := 1; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if isFirstColumnZero {
		for i := 1; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
