package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	var pos = map[int]int{}
	for i, num := range nums {
		if p, ok := pos[num]; ok && i-p <= k {
			return true
		}
		pos[num] = i
	}

	return false
}
