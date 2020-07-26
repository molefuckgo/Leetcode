package main

import "fmt"

func main() {
	var matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		// {5, 6},
	}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	numberList := make([]int, 0)
	matrixRows := len(matrix)
	if len(matrix) == 0 {
		return numberList
	}
	matrixColumns := len(matrix[0])
	l := 0
	r := matrixColumns - 1
	t := 0
	b := matrixRows - 1
	for l <= r && t <= b {
		for i := l; i <= r; i++ {
			numberList = append(numberList, matrix[t][i])
		}
		t++
		if t > b {
			return numberList
		}

		for i := t; i <= b; i++ {
			numberList = append(numberList, matrix[i][r])
		}
		r--
		if l > r {
			return numberList
		}

		for i := r; i >= l; i-- {
			numberList = append(numberList, matrix[b][i])
		}
		b--
		if b < t {
			return numberList
		}

		for i := b; i >= t; i-- {
			numberList = append(numberList, matrix[i][l])
		}
		l++

	}
	return numberList
}
