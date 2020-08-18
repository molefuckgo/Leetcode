package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//a := []int{1, 2, 3}
	//b := []int{1, 2, 3}
	//fmt.Println(a == b)
}

func findUnsortedSubarray(nums []int) int {
	numsTemp := make([]int, len(nums))
	copy(numsTemp, nums)
	sortArray(nums)
	var start, end int
	fmt.Println(numsTemp)
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if numsTemp[i] != nums[i] {
			start = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if numsTemp[i] != nums[i] {
			end = i
			break
		}
	}
	fmt.Println(start, end)
	return end - start + 1
}

func randomizedQuicksort(nums *[]int, l, r int) {
	if r-l <= 0 {
		return
	}
	mid := randomizedPartion(nums, l, r)
	randomizedQuicksort(nums, l, mid-1)
	randomizedQuicksort(nums, mid+1, r)
}

func sortArray(nums []int) []int {
	randomizedQuicksort(&nums, 0, len(nums)-1)
	return nums
}

func randomizedPartion(nums *[]int, l, r int) int { // 对l, r进行划分，返回分界下标pos
	pivot := rand.Intn(r-l) + l
	(*nums)[r], (*nums)[pivot] = (*nums)[pivot], (*nums)[r]
	i := l - 1
	for j := l; j < r; j++ {
		if (*nums)[j] < (*nums)[r] {
			i++
			(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
		}
	}
	i++
	(*nums)[i], (*nums)[r] = (*nums)[r], (*nums)[i]
	return i
}
