package main

func main() {
	//[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
}

var rows, columns int

func solve(board [][]byte) {
	rows = len(board)
	if rows == 0 {
		return
	}
	columns = len(board[0])

	for i := 0; i < rows; i++ {
		dfs(&board, i, 0)
		dfs(&board, i, columns-1)
	}

	for i := 0; i < columns; i++ {
		dfs(&board, 0, i)
		dfs(&board, rows-1, i)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board *[][]byte, row, column int) {
	if row < 0 || row >= rows || column < 0 || column >= columns || (*board)[row][column] != 'O' {
		return
	}
	(*board)[row][column] = 'A'
	dfs(board, row+1, column)
	dfs(board, row-1, column)
	dfs(board, row, column+1)
	dfs(board, row, column-1)
}
