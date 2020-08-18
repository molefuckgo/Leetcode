package main

import "fmt"

func main() {
	nums := [][]int{
		{1, 2, 3, 4},
	}
	r := 2
	c := 2
	fmt.Println(matrixReshape(nums, r, c))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	rows := len(nums)
	if rows == 0 {
		return nums
	}
	columns := len(nums[0])
	if columns == 0 {
		return nums
	}
	if rows*columns != r*c {
		return nums
	}
	results := make([][]int, r)
	row := 0
	column := 0
	for i := 0; i < r; i++ {
		results[i] = make([]int, c)
		for j := 0; j < c; j++ {
			results[i][j] = nums[row][column]
			row, column = getNextIJ(row, column, rows, columns)
		}
	}
	return results
}

func getNextIJ(row, column, rows, columns int) (int, int) {
	if column+1 == columns {
		column = 0
		row++
		return row, column
	}
	column++
	return row, column
}
