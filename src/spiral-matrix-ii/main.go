package main

import "fmt"

func main() {
	n := 3
	fmt.Println(generateMatrix(n))
}

func generateMatrix(n int) [][]int {
	l := 0
	r := n - 1
	t := 0
	b := n - 1
	mat := make([][]int, n, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}
	num := 1
	tar := n * n
	for num <= tar {
		for i := l; i <= r; i++ {
			mat[t][i] = num
			num++
		}
		t++ // 上边界+1

		for i := t; i <= b; i++ {

			mat[i][r] = num
			num++
		}
		r-- // 右边界-1

		for i := r; i >= l; i-- {
			mat[b][i] = num
			num++
		}
		b--

		for i := b; i >= t; i-- {
			mat[i][l] = num
			num++
		}
		l++

	}
	return mat
}

func spiralMatrix(begin int, n int, t int, matrixArray [][]int) [][]int {
	i := 0
	j := 0
	for j < n {
		matrixArray[i+t][j+t] = begin
		begin++
		if j+1 < n {
			j++
		} else {
			break
		}
	}
	i++
	for i < n {
		matrixArray[i+t][j+t] = begin
		begin++
		if i+1 < n {
			i++
		} else {
			break
		}
	}
	j--
	for j >= 0 {
		matrixArray[i+t][j+t] = begin
		begin++
		if j-1 >= 0 {
			j--
		} else {
			break
		}
	}
	i--
	for ; i >= 1; i-- {
		matrixArray[i+t][j+t] = begin
		begin++
	}
	if n-2 > 0 {
		t++
		n -= 2
		return spiralMatrix(begin, n, t, matrixArray)
	} else {
		return matrixArray
	}

}
