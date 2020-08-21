package main

import "math"

func main() {

}

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	num3 := n / 3
	remaining := n % 3
	result := int(math.Pow(float64(3), float64(num3)))
	if remaining == 2 {
		result *= 2
	} else if remaining == 1 {
		result = result / 3 * 4
	}
	return result
}
