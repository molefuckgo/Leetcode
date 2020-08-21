package main

import "fmt"

func main() {
	fmt.Println(numWays(3, 2))
}

//输入: n = 3，k = 2
//输出: 6
//解析: 用 c1 表示颜色 1，c2 表示颜色 2，所有可能的涂色方案有:
//
//            柱 1    柱 2   柱 3
// -----      -----  -----  -----
//   1         c1     c1     c2 	第1个柱子k种选择，第二个柱子k种选择，第三个有分支，如果第二个选择了和第一个相同的颜色，则有k - 1种选择，否则是k种
//   2         c1     c2     c1
//   3         c1     c2     c2
//   4         c2     c1     c1
//   5         c2     c1     c2
//   6         c2     c2     c1

func numWays(n int, k int) int {
	mempArr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i == 1 {
			mempArr[i] = k
		} else if i == 2 {
			mempArr[2] = k * k
		} else {
			mempArr[i] = mempArr[i-1]*(k-1) + mempArr[i-2]*(k-1)
		}
	}
	return mempArr[n]
}
