package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			break
		}
	}
	if digits[0] == 0 {
		digits = make([]int, len(digits)+1)
		digits[0] = 1
		digits = digits
	}
	return digits
}
