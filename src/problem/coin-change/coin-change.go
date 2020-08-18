package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5, 8}, 11))
}

var memoMap map[int]int

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	memoMap = make(map[int]int)
	for i := 0; i < len(coins); i++ {
		if amount == coins[i] {
			return 1
		}
		memoMap[coins[i]] = 1
	}
	return dp(coins, amount)
}

func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if minNum, ok := memoMap[amount]; ok {
		return minNum
	}
	result := math.MaxInt64
	for i := 0; i < len(coins); i++ {
		subProblem := dp(coins, amount-coins[i])
		if subProblem == -1 {
			continue
		}
		result = min2(result, subProblem+1)
	}
	if result == math.MaxInt64 {
		result = -1
	}
	memoMap[amount] = result
	return result
}

func min2(x, y int) int {
	if x >= y {
		return y
	}
	return x
}
