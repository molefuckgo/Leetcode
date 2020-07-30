package main

import "fmt"

func main() {
	var nums = []int{2, 2, 2, 3, 4, 5}
	fmt.Println(findMagicIndex(nums))
}

func findMagicIndex(nums []int) int {
	return getAnswer(nums, 0, len(nums)-1)
}

func getAnswer(nums []int, left, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	leftAnswer := getAnswer(nums, left, mid-1)
	if leftAnswer != -1 {
		return leftAnswer
	} else if mid == nums[mid] {
		return mid
	}
	rightAnswer := getAnswer(nums, mid+1, right)
	if rightAnswer != -1 {
		return rightAnswer
	}
	return -1
}
