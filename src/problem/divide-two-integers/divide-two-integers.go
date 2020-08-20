package main

import (
	"fmt"
	"math"
)

func main() {
	//1004958205
	//-2137325331
	fmt.Println(divide(3, -66))
	//fmt.Println(math.MinInt32)
	//fmt.Println(math.MaxInt32)
}

func divide(dividend int, divisor int) int {
	positiveOrNegative := dividend^divisor >= 0 // 正负
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))
	result := 1
	nowNum := divisor
	for nowNum <= dividend {
		nowNum = nowNum << 1
		result = result << 1
	}
	nowNum = nowNum >> 1
	result = result >> 1
	//if nowNum == dividend {
	//	return result
	//}

	for nowNum < dividend {
		nowNum += divisor
		result++
	}
	if nowNum > dividend {
		result--
	}
	if dividend < divisor {
		result = 0
	}

	if !positiveOrNegative {
		result = -result
	}

	if result >= math.MaxInt32 {
		result = math.MaxInt32
	} else if result <= math.MinInt32 {
		result = math.MinInt32
	}

	return result
}
