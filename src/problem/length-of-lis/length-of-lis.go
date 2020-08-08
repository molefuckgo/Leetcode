package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7}))
}

func lengthOfLIS(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	dp := make([]int, numsLen)
	dp[0] = 1
	maxLength := dp[0]
	for i := 1; i < numsLen; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] == 0 {
			dp[i] = 1
		}
		maxLength = max2(maxLength, dp[i])
	}

	return maxLength
}

func max2(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
