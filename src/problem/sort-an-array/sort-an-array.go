package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	nums := []int{5, 1, 1, 2, 0, 0}
	sort.Ints(nums)
	fmt.Println(nums)
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
