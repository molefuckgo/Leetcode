package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
}

func titleToNumber(s string) int {
	sLen := len(s)
	result := 0
	for i := sLen - 1; i >= 0; i-- {
		result += int(s[i]-'@') * intPow(26, sLen-i-1)
	}
	return result
}

func intPow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func findNumberOfLIS(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	dp := make([][]int, numsLen)
	for i := 0; i < numsLen; i++ {
		dp[i] = []int{1, 1}
	}
	dp[0] = []int{1, 1} // 长度是1，路数是1
	for i := 1; i < numsLen; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i][0] < dp[j][0]+1 {
					dp[i][0] = dp[j][0] + 1
					dp[i][1] = dp[j][1]
				} else if dp[i][0] == dp[j][0]+1 {
					dp[i][1] += dp[j][1]
				}
			}
		}
	}
	result := 0
	maxLength := 0
	for i := 0; i < numsLen; i++ {
		if maxLength < dp[i][0] {
			maxLength = dp[i][0]
			result = dp[i][1]
		} else if maxLength == dp[i][0] {
			result += dp[i][1]
		}
	}
	return result
}
