package main

import "fmt"

func main() {
	fmt.Println(convert("LEETCODEISHIRING", 4))
}

func convert(s string, numRows int) string {
	sLen := len(s)
	if numRows == 1 {
		return s
	} else if numRows < 1 || sLen == 0 {
		return ""
	}

	afterS := ""
	index := 0
	numRow := 0
	var middleIndex []int
	for i := 0; i < sLen; i++ {

		if len(middleIndex) == 0 {
			middleIndex = getMiddleIndex(numRows, numRow, index)
		}

		if index < sLen {
			afterS += string(s[index])
			index += middleIndex[0]
			middleIndex = middleIndex[1:]
		} else {
			index = numRow + 1
			numRow++
			middleIndex = make([]int, 0)
			i--
		}
	}
	return afterS
}

func getMiddleIndex(numRows int, numRow int, index int) []int { //根据总行数和现在的行数推断出中间的index
	if numRow == 0 || numRow == numRows-1 {
		return []int{2 * (numRows - 1), 2 * (numRows - 1)}
	} else {
		return []int{(numRows - numRow - 1) * 2, 2 * numRow}
	}
}
