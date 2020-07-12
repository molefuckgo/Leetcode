package main

import "fmt"

func main() {
	fmt.Println(MaxB(7))
}

func maxA(N int) int {

	return dp(N, 0, 0)
}
func dp(n int, aNum int, copyNum int) int {
	if n <= 0 {
		return aNum
	}
	return max3(
		dp(n-1, aNum+1, copyNum),       // A
		dp(n-1, aNum+copyNum, copyNum), // CTRL-V
		dp(n-2, aNum, aNum),            // CTRL-A CTRL-C
	)
	return 0
}

func MaxB(N int) int {
	dp := make([]int, N+1)
	for i := 1; i <= N; i++ {
		// 按A键
		dp[i] = dp[i-1] + 1
		for j := 2; j < i; j++ {
			// 全屏复制d[i - 2]，连续粘贴i - j次
			dp[i] = max2(dp[i], dp[j-2]*(i-j+1))
		}
	}
	// dp[1] = 3
	// fmt.Println(dp[1])
	// return 0
	return dp[N]
}

func max3(a int, b int, c int) int {

	maxNum := 0
	if a >= b && a >= c {
		maxNum = a
	} else if b >= a && b >= c {
		maxNum = b
	} else {
		maxNum = c
	}
	return maxNum
}

func max2(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
