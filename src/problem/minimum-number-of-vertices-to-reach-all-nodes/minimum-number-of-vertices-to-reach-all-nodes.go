package main

import "fmt"

func main() {
	fmt.Println(findSmallestSetOfVertices(6, [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}))
}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	var result []int
	var resMap = make(map[int]bool)
	for _, curRes := range edges {
		resMap[curRes[1]] = true
	}
	for i := 0; i < n; i++ {
		if resMap[i] == false {
			result = append(result, i)
		}
	}
	return result
}
