package main

import "fmt"

func main() {
	fmt.Println(findSubsequences([]int{3, 8, 7, 7, 6}))
}

func findSubsequences(nums []int) [][]int {
	memoArr := make([][]int, 0)
	for _, num := range nums {
		hasLessFlag := false
		for i := 0; i < len(memoArr); i++ {
			if memoArr[i][len(memoArr[i])-1] <= num {
				memoArr[i] = append(memoArr[i], num)
				hasLessFlag = true
			}

		}
		if !hasLessFlag {

		}
	}
	fmt.Println(memoArr)
	return memoArr
}
