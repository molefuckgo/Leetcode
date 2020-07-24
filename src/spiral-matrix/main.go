package main

import "fmt"

func main() {
	var matrix = [][]int{
		{2, 3},
		// {4, 5, 6},
		// {7, 8, 9},
		// {10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	matrixRows := len(matrix)

	if matrixRows == 0 {
		return make([]int, 0)
	}
	matrixColumns := len(matrix[0])
	result := make([]int, 0)
	return spiralMatrix(matrixRows, matrixColumns, 0, matrix, result)
}

func spiralMatrix(matrixRows int, matrixColumns int, t int, matrix [][]int, result []int) []int {
	i := 0
	j := 0
	for j < matrixColumns {
		result = append(result, matrix[i+t][j+t])
		if j+1 < matrixColumns {
			j++
		} else {
			break
		}
	}
	i++
	for i < matrixRows {
		result = append(result, matrix[i+t][j+t])
		if i+1 < matrixRows {
			i++
		} else {
			break
		}
	}
	j--
	for j >= 0 && j+t < matrixColumns-1 {
		result = append(result, matrix[i+t][j+t])
		if j-1 >= 0 {
			j--
		} else {
			break
		}
	}
	i--
	for ; i >= 1; i-- {
		result = append(result, matrix[i+t][j+t])
	}
	if matrixColumns-2 > 0 {
		t++
		return spiralMatrix(matrixRows-2, matrixColumns-2, t, matrix, result)
	} else {
		return result
	}

}
