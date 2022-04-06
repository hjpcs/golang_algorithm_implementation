package is_valid_sudoku

func isValidSudoku(board [][]byte) bool {
	return isRowValid(board) && isColumnValid(board) && isGridValid(board)
}

func isRowValid(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		resultMap := make(map[byte]int)
		for _, v := range board[i] {
			resultMap[v]++
		}
		for k, v := range resultMap {
			if k != '.' && v > 1 {
				return false
			}
		}
	}
	return true
}

func isColumnValid(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		resultMap := make(map[byte]int)
		for j := 0; j < 9; j++ {
			resultMap[board[j][i]]++
		}
		for k, v := range resultMap {
			if k != '.' && v > 1 {
				return false
			}
		}
	}
	return true
}

func isGridValid(board [][]byte) bool {
	for i := 1; i <= 7; i += 3 {
		for j := 1; j <= 7; j += 3 {
			resultMap := make(map[byte]int)
			resultMap[board[i-1][j-1]]++
			resultMap[board[i-1][j]]++
			resultMap[board[i-1][j+1]]++
			resultMap[board[i][j-1]]++
			resultMap[board[i][j]]++
			resultMap[board[i][j+1]]++
			resultMap[board[i+1][j-1]]++
			resultMap[board[i+1][j]]++
			resultMap[board[i+1][j+1]]++
			for k, v := range resultMap {
				if k != '.' && v > 1 {
					return false
				}
			}
		}
	}
	return true
}
