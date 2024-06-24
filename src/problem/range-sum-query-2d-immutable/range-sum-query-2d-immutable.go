package main

import "fmt"

func main() {
	var numMatrix = Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	fmt.Println(numMatrix.SumRegion(2, 1, 4, 3))
}

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var m = len(matrix)
	if m == 0 {
		return NumMatrix{}
	}
	var n = len(matrix[0])
	if n == 0 {
		return NumMatrix{}
	}
	var sums = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		sums[i] = make([]int, n+1)
	}
	for i := 1; i < len(sums); i++ {
		for j := 1; j < len(sums[0]); j++ {
			sums[i][j] = matrix[i-1][j-1] + sums[i][j-1] + sums[i-1][j] - sums[i-1][j-1]
		}
	}
	return NumMatrix{
		sums: sums,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var result = this.sums[row2+1][col2+1] - this.sums[row1][col2+1] - this.sums[row2+1][col1] + this.sums[row1][col1]
	return result
}
