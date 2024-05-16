package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	var result = [][]int{{1}}

	for i := 1; i < numRows; i++ {
		result = append(result, generateOne(i+1, result[i-1]))
	}
	return result
}

func generateOne(numRows int, before []int) []int {
	var oneResult = []int{}
	for i := 0; i < numRows; i++ {
		if i == 0 || len(before) == i {
			oneResult = append(oneResult, before[len(before)-1])
		} else {
			oneResult = append(oneResult, before[i-1]+before[i])
		}
	}
	return oneResult
}
