package main

func main() {

}

func fib(N int) int {
	memoMap := make(map[int]int)
	return helper(memoMap, N)
}

func helper(memoMap map[int]int, N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	if memoMap[N] != 0 {
		return memoMap[N]
	}
	memoMap[N] = helper(memoMap, N-1) + helper(memoMap, N-2)
	return memoMap[N]
}
