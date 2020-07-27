package main

import "fmt"

func main() {
	fmt.Println(reverse(1534236469))
	//result := 2
	//for i := 1;i < 31;i++ {
	//	result *= 2
	//}
	//fmt.Println(result)
}

func reverse(x int) int {
	var positiveorNegative bool
	if x < 0 {
		x = -x
		positiveorNegative = false
	} else {
		positiveorNegative = true
	}
	xReverse := 0
	for ; x > 0; x /= 10 {
		xReverse = (10 * xReverse) + x%10
	}
	if positiveorNegative {
		if xReverse <= 2147483647 {
			return xReverse
		}
		return 0
	} else {
		if xReverse <= 2147483648 {
			return -xReverse
		}
		return 0
	}
}
