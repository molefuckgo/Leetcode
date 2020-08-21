package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, -2}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < numsLen; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]

		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
