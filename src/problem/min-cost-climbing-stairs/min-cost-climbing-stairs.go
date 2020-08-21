package main

import "fmt"

func main() {
	//fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}

func minCostClimbingStairs(cost []int) int {
	costLen := len(cost)
	memoArr := make([]int, costLen)
	memoArr[0] = cost[0]
	memoArr[1] = cost[1]
	for i := 2; i < costLen; i++ {
		memoArr[i] = cost[i] + min2(memoArr[i-1], memoArr[i-2])
	}
	return min2(memoArr[costLen-1], memoArr[costLen-2])
}

func min2(x, y int) int {
	if x >= y {
		return y
	}
	return x
}
