package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 2, 2, 2, 4, 4}))
}

func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	var lastNum int
	i := 0
	for i < len(nums) {
		if i == 0 && len(nums) != 0 {
			lastNum = nums[i]
			i++
		} else {
			if nums[i] == lastNum {
				numsLen--
				deleteArrayIndex(&nums, i)
			} else {
				lastNum = nums[i]
				i++
			}
		}

	}
	fmt.Println(nums)
	return numsLen
}

func deleteArrayIndex(nums *[]int, index int) {
	*nums = append((*nums)[:index], (*nums)[index+1:]...)
}
