package main

import "fmt"

func main() {
	var nums = []int{0, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	removeValElement(&nums, 0)
	for i := len(nums); i < cap(nums); i++ {
		nums = append(nums, 0)
	}
}

func removeValElement(nums *[]int, val int) {
	var slow, fast = 0, 0
	for fast < len(*nums) {
		if (*nums)[fast] != val {
			(*nums)[slow] = (*nums)[fast]
			slow++
		}
		fast++
	}
	*nums = (*nums)[:slow]
}

func deleteArrayIndex(nums *[]int, index int) {
	*nums = append((*nums)[:index], (*nums)[index+1:]...)
}
