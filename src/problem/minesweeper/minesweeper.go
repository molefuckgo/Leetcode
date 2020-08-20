package main

import "fmt"

func main() {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}
	updateBoard(board, []int{1, 2})
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Print(string(board[i][j]), " ")
		}
		fmt.Println()
	}
}

func updateBoard(board [][]byte, click []int) [][]byte {
	rows := len(board)
	if rows == 0 {
		return board
	}
	columns := len(board[0])
	if columns == 0 {
		return board
	}
	dfs(&board, click, rows, columns)
	return board
}

func dfs(board *[][]byte, click []int, rows, columns int) { // 最后返回有几个雷

	if (*board)[click[0]][click[1]] == 'B' { // 已经翻过了
		return
	}
	if (*board)[click[0]][click[1]] == 'M' { // 是雷变成X直接返回雷数目
		(*board)[click[0]][click[1]] = 'X'
		return
	}
	if ((*board)[click[0]][click[1]] >= '1' && (*board)[click[0]][click[1]] <= '9') || (*board)[click[0]][click[1]] == 'X' {
		if (*board)[click[0]][click[1]] == 'X' {
			return
		}
		return
	}
	thisXNum := 0
	for i := click[0] - 1; i <= click[0]+1; i++ {
		if i < 0 || i >= rows {
			continue
		}
		for j := click[1] - 1; j <= click[1]+1; j++ {
			if j < 0 || j >= columns {
				continue
			}
			if (*board)[i][j] == 'M' || (*board)[i][j] == 'X' {
				thisXNum++
			}
		}
	}
	if thisXNum == 0 {
		(*board)[click[0]][click[1]] = 'B'
	} else {
		(*board)[click[0]][click[1]] = byte(thisXNum) + '0'
		return
	}
	if click[0] > 0 {
		//fmt.Println("click[0] > 0")
		dfs(board, []int{click[0] - 1, click[1]}, rows, columns)
	}
	if click[0] < rows-1 {
		//fmt.Println("click[0] < rows-1")
		dfs(board, []int{click[0] + 1, click[1]}, rows, columns)
	}
	if click[1] > 0 {
		//fmt.Println("click[1] > 0")
		dfs(board, []int{click[0], click[1] - 1}, rows, columns)
	}
	if click[1] < columns-1 {
		//fmt.Println("click[1] < columns-1")
		dfs(board, []int{click[0], click[1] + 1}, rows, columns)
	}

	if click[0] > 0 && click[1] > 0 {
		dfs(board, []int{click[0] - 1, click[1] - 1}, rows, columns)
	}
	if click[0] < rows-1 && click[1] < columns-1 {
		dfs(board, []int{click[0] + 1, click[1] + 1}, rows, columns)
	}
	if click[0] > 0 && click[1] < columns-1 {
		dfs(board, []int{click[0] - 1, click[1] + 1}, rows, columns)
	}
	if click[0] < rows-1 && click[1] > 0 {
		dfs(board, []int{click[0] + 1, click[1] - 1}, rows, columns)
	}

}
