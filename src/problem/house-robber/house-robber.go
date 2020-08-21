package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	if numsLen == 1 {
		return nums[0]
	}
	max := nums[0]

	for i := 1; i < numsLen; i++ {
		if i == 1 {
			nums[i] = max2(nums[i], nums[i-1])
		} else {
			nums[i] = max2(nums[i]+nums[i-2], nums[i-1])
		}
		max = max2(max, nums[i])
	}
	return max
}

func max2(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
