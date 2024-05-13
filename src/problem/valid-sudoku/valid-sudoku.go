package main

import "fmt"

func main() {
	var board [][]byte = [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}}
	fmt.Println(isValidSudoku(board))
}
func isValidSudoku(board [][]byte) bool {
	return checkAllLine(board) && checkAllList(board)
}

func checkAllLine(board [][]byte) bool {
	for _, oneBoard := range board {
		if !checkOneLine(oneBoard) {
			return false
		}
	}
	return true
}

func checkOneLine(oneLine []byte) bool {
	var recordMap = make(map[byte]bool)
	for _, v := range oneLine {
		if recordMap[v] == true && v != '.' {
			return false
		} else {
			recordMap[v] = true
		}
	}
	return true
}

func tranOneListToLineAndCheck(board [][]byte, i int) bool {
	var resultLine = make([]byte, len(board))
	for j, oneLine := range board {
		resultLine[j] = oneLine[i]
	}
	return checkOneLine(resultLine)
}

func checkAllList(board [][]byte) bool {
	for i, _ := range board {
		if !tranOneListToLineAndCheck(board, i) {
			return false
		}
	}
	return true
}
