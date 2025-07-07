package backtracking

func Exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[0] {
			if dfs5(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs5(board [][]byte, i, j, index int, word string) bool {
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) {
		return false
	}
	if board[i][j] != word[index] || board[i][j] == '-' {
		return false
	}
	if index == len(word)-1 {
		return true
	}
	c := board[i][j]
	board[i][j] = '-'

	found := dfs5(board, i+1, j, index+1, word) ||
		dfs5(board, i, j-1, index+1, word) ||
		dfs5(board, i, j+1, index+1, word) ||
		dfs5(board, i-1, j, index+1, word)

	board[i][j] = c
	return found

}
