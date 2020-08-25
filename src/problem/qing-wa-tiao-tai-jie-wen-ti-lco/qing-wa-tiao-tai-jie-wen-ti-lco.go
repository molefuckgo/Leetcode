package main

func main() {

}

func numWays(n int) int {
	memoMap := make([]int, n+1)
	return dfs(n, memoMap) % (1e9 + 7)
}

func dfs(n int, memoMap []int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		memoMap[n] = n
		return n
	}
	var firstResult, secondResult int
	if memoMap[n-2] > 0 {
		firstResult = memoMap[n-2]
	} else {
		firstResult = dfs(n-2, memoMap)
		memoMap[n-2] = firstResult
	}

	if memoMap[n-1] > 0 {
		secondResult = memoMap[n-1]
	} else {
		secondResult = dfs(n-1, memoMap)
		memoMap[n-1] = secondResult
	}
	return (firstResult + secondResult) % (1e9 + 7)
}
