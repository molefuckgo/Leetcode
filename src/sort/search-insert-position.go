package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 4))
}

func searchInsert(nums []int, target int) int {
	numsLen := len(nums)
	start := 0
	end := numsLen - 1
	mid := (start + end) / 2
	for start < end {
		if nums[mid] == target {
			for mid >= 1 && nums[mid-1] == nums[mid] {
				mid--
			}
			return mid
		} else if nums[mid] > target {
			if mid == 0 {
				return mid
			} else {
				end = mid - 1
				mid = (start + end) / 2
			}
		} else {
			if mid == numsLen-1 {
				return mid + 1
			} else {
				start = mid + 1
				mid = (start + end) / 2
			}
		}
	}
	if nums[mid] < target {
		return mid + 1
	}
	return mid
}
