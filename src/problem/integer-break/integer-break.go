package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(integerBreak(15))
}

func integerBreak(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	} else if n == 4 {
		return 4
	} else if n == 5 {
		return 6
	} else if n%3 == 0 {
		return int(math.Pow(float64(3), float64(n/3)))
	}
	timesThree := n/3 - 1
	beforeResult := int(math.Pow(float64(3), float64(timesThree)))
	remanent := n - 3*timesThree
	if remanent == 2 || remanent == 5 {
		return beforeResult * 6
	} else if remanent == 3 {
		return beforeResult * 3
	} else {
		return beforeResult * 4
	}
}
