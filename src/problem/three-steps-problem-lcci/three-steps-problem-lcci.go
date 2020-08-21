package main

import "fmt"

func main() {
	fmt.Println(waysToStep(1000))
}

func waysToStep2(n int) int {
	memoArr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i == 1 {
			memoArr[i] = 1
		} else if i == 2 {
			memoArr[i] = 2
		} else if i == 3 {
			memoArr[i] = 4
		} else {
			memoArr[i] = (memoArr[i-1] + memoArr[i-2] + memoArr[i-3]) % 1000000007
		}
	}
	return memoArr[n]
}

func waysToStep(n int) int {
	r1 := 1
	if n == 1 {
		return r1
	}
	r2 := 2
	if n == 2 {
		return r2
	}
	r3 := 4
	if n == 3 {
		return r3
	}
	for i := 4; i <= n; i++ {
		r3Temp := r3
		r3 = (r1 + r2 + r3) % 1000000007
		r1 = r2
		r2 = r3Temp
	}
	return r3
}
