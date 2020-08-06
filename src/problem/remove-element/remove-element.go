package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{1, 2, 2, 3, 3, 4, 4}, 3))
}

func removeElement(nums []int, val int) int {
	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}
	return i
}
