package main

import "fmt"

func main() {
	fmt.Println(maxValue([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
		{3, 7, 9},
	}))
}

func maxValue(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	columns := len(grid[0])
	if columns == 0 {
		return 0
	}
	for i := 1; i < rows; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < columns; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			grid[i][j] += max2(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[rows-1][columns-1]
}

func max2(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
