package main

import "fmt"

func main() {
	mmap := make(map[int]int)
	length, ok := mmap[4]
	fmt.Println(length, ok)
}

func longestSubsequence(arr []int, difference int) int {
	arrLen := len(arr)
	if arrLen == 0 {
		return 0
	}
	maxLength := 1
	memoMap := make(map[int]int)
	memoMap[arr[0]] = 1
	for i := 1; i < arrLen; i++ {
		length := memoMap[arr[i]-difference]
		memoMap[arr[i]] = length + 1

		maxLength = max2(memoMap[arr[i]], maxLength)
	}
	return maxLength
}

func max2(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
