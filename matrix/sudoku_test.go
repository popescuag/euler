package matrix

func IsValidSudoku(board [][]byte) bool {
	rowMap := make(map[byte]bool)
	colMap := make(map[byte]bool)

	nr, nc := len(board), len(board[0])

	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if board[i][i] == '.' {
				continue
			}
			if _, ok := rowMap[board[i][j]]; ok {
				return false
			}
			if _, ok := colMap[board[i][j]]; ok {
				return false
			}
			rowMap[byte(i)] = true
			colMap[byte(j)] = true
		}
	}

	return IsValidSudoku(board[:3][:3]) && IsValidSudoku(board[:3][3:6]) && IsValidSudoku(board[:3][6:9]) &&
		IsValidSudoku(board[3:6][:3]) && IsValidSudoku(board[3:9][3:6]) && IsValidSudoku(board[3:6][6:9]) &&
		IsValidSudoku(board[6:9][:3]) && IsValidSudoku(board[6:9][3:6]) && IsValidSudoku(board[6:9][6:9])
}
