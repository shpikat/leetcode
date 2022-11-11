package leetcode

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	// Real quick check
	if len(word) > m*n {
		return false
	}

	// A bit slower, but still quicker than the full check
	letters := [128]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			letters[board[i][j]]++
		}
	}
	for _, r := range word {
		if letters[r] == 0 {
			return false
		}
		letters[r]--
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if findWord(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func findWord(board [][]byte, row int, column int, word string, index int) bool {
	exists := false
	if board[row][column] == word[index] {
		board[row][column] = 0
		next := index + 1
		exists = next == len(word) ||
			(column > 0 && findWord(board, row, column-1, word, next)) ||
			(column < len(board[row])-1 && findWord(board, row, column+1, word, next)) ||
			(row > 0 && findWord(board, row-1, column, word, next)) ||
			(row < len(board)-1 && findWord(board, row+1, column, word, next))

		board[row][column] = word[index]
	}

	return exists
}
