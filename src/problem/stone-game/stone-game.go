package main

import "fmt"

func main() {
	fmt.Println(stoneGame([]int{3, 9, 1, 2}))
}

const (
	first  = 0 // 先手
	second = 1 // 后手
)

func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][][]int, n)
	//for i := 0;i < n;i++ {
	//	dp[i][i] = []int{piles[i], 0}
	//}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i] = append(dp[i], []int{0, 0})
		}
	}
	fmt.Println(dp)
	for i := 0; i < n; i++ {
		for j := 0; j+i < n; j++ {
			if i == 0 {
				dp[j][i+j] = []int{piles[j], 0}
			} else {
				fmt.Println(i + j + 1)
				dp[j][i+j][first] = max2(piles[j]+dp[j+1][i+j][second], piles[i+j]+dp[j][i+j-1][second])
				dp[j][i+j][second] = max2(piles[j]+dp[j+1][i+j][first], piles[i+j]+dp[j][i+j-1][first])
			}

		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(dp[i])
	}
	return false
}
func max2(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
