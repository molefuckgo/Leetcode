package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(massage(nums))
}

func massage(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	memoArr := make([]int, numsLen)
	memoArr[0] = nums[0]
	max := memoArr[0]
	for i := 1; i < numsLen; i++ {
		if i == 1 {
			memoArr[i] = max2(nums[1], memoArr[0])
		} else {
			memoArr[i] = max2(nums[i]+memoArr[i-2], memoArr[i-1])
		}
		max = max2(memoArr[i], max)
	}
	return max
}
func max2(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
