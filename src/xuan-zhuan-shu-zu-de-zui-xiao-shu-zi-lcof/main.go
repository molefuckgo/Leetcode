package main

import "fmt"

func main() {
	var arr = []int{2, 2, 2, 0, 1}
	fmt.Println(minArray(arr))
}

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		middle := (left + right) / 2
		if numbers[right] < numbers[middle] {
			left = middle + 1
		} else if numbers[right] > numbers[middle] {
			right = middle - 1
		} else { // 相等的情况缩小右边界，这样至少中间最小的能留下，不会影响结果
			right -= 1
		}
	}
	return numbers[right]
}
