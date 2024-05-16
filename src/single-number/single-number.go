package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 2, 6, 9, 9}))
}

func singleNumber(nums []int) int {
	var single = 0

	for _, value := range nums {
		single ^= value
	}
	return single
}
