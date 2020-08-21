package main

import (
	"fmt"
)

func main() {
	fmt.Println(cuttingRope(120))
}

func cuttingRope(n int) int {
	if n == 3 {
		return 2
	}
	nums3 := n/3 - 1
	var result, remaining int
	result = 1
	if nums3 > 0 {
		//result = int(math.Pow(float64(3), float64(nums3)))
		for i := 0; i < nums3; i++ {
			result = result * 3 % 1000000007
		}
		remaining = n - 3*nums3
	} else {
		remaining = n
	}

	// 剩余

	if remaining == 2 {
		result = result % 1000000007
	} else if remaining == 3 { // 刚好除完
		result = result * 3 % 1000000007
	} else if remaining == 4 { // 还剩4
		result = result * 4 % 1000000007
	} else { //
		result = result * 2 * 3 % 1000000007
	}
	return result
}
