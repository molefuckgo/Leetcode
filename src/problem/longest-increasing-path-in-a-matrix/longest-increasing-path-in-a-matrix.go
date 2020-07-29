package main

import "fmt"

func main() {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
		//{3, 4, 5},
		//{3, 2, 6},
		//{2, 2, 1},
		//{1, 3},
		//{2, 4},
	}
	fmt.Println(longestIncreasingPath(matrix))
}

var (
	directions               = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
	rows, columns, maxLength int
)

func longestIncreasingPath(matrix [][]int) int {
	rows = len(matrix)
	maxLength = 0

	if rows == 0 {
		return maxLength
	}

	columns = len(matrix[0])
	if columns == 0 {
		return maxLength
	}

	memo := make([][]int, rows)
	for i := 0; i < rows; i++ {
		memo[i] = make([]int, columns)
	}
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if memo[row][column] == 0 {
				memo = depthFirstSearch(matrix, memo, row, column)
			}
			maxLength = max(maxLength, memo[row][column])
		}
	}
	return maxLength
}

func depthFirstSearch(matrix [][]int, memo [][]int, row int, column int) [][]int {
	if memo[row][column] != 0 {
		return memo
	}
	for _, direction := range directions {
		newRow := row + direction[0]
		newColumn := column + direction[1]
		if check(newRow, newColumn) && matrix[newRow][newColumn] > matrix[row][column] { // 在范围内
			memo = depthFirstSearch(matrix, memo, newRow, newColumn)
			if memo[row][column] < memo[newRow][newColumn]+1 {
				memo[row][column] = memo[newRow][newColumn] + 1
			}
		}

	}
	if memo[row][column] == 0 {
		memo[row][column] = 1
	}
	return memo
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func check(newRow, newColumn int) bool {
	if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns {
		return true
	}
	return false
}
