package main

import "fmt"

func main() {
	fmt.Println(longestSubsequence([]int{1, 3, 5, 6, 7, 8, 9, 10}, 2))
}

func longestSubsequence(arr []int, difference int) int {
	arrLen := len(arr)
	if arrLen == 0 {
		return 0
	}
	memo := make(map[int]bool)
	maxLength := 1
	dp := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		if memo[i] {
			continue
		}
		memo[i] = true
		lastNum := arr[i]
		dp[i]++

		for j := i + 1; j < arrLen; j++ {
			if !memo[j] && arr[j]-lastNum == difference {
				memo[j] = true
				lastNum = arr[j]
				dp[i]++
			}
		}
		maxLength = max2(maxLength, dp[i])
		if arrLen-i <= maxLength {
			return maxLength
		}
	}
	return maxLength
}

func max2(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
