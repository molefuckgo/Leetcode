package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrt(8))
}

func mySqrt(x int) int {
	var y = float64(x)
	for math.Abs(float64(y*y-float64(x))) > 1e-6 {
		y = (y + float64(x)/y) / 2
	}
	return int(y)
}
