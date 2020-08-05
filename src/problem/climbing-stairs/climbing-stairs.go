package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	memo := make([]int, n+1)
	return dfs(n, memo)
}

func dfs(n int, memo []int) int {
	if n <= 2 {
		memo[n] = n
		return n
	}
	var firstResult, secondResult int
	if memo[n-1] > 0 {
		firstResult = memo[n-1]
	} else {
		firstResult = dfs(n-1, memo)
		memo[n-1] = firstResult
	}
	if memo[n-2] > 0 {
		secondResult = memo[n-2]
	} else {
		secondResult = dfs(n-2, memo)
		memo[n-2] = secondResult
	}

	return firstResult + secondResult
}
