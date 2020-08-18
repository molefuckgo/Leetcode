package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(quickSort(nums, 0, len(nums)-1))
}

//func randomizedPartion(nums *[]int, l, r int) int {	// 对l, r进行划分，返回分界下标pos
//	pivot := rand.Intn(r - l) + l	// 随机分界值
//	//pivot := 2	// 随机分界值
//	(*nums)[pivot], (*nums)[r] = (*nums)[r], (*nums)[pivot]
//	i := l - 1
//	for j := l;j < r;j++ {
//		if (*nums)[j] < (*nums)[r] {
//			i++
//			(*nums)[j], (*nums)[i] = (*nums)[i], (*nums)[j]
//		}
//	}
//	i++
//	(*nums)[r], (*nums)[i] = (*nums)[i], (*nums)[r]
//	return i
//}
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

//func randomizedQuicksort(nums *[]int, l, r int ) {
//	if r - l <= 0 {
//		return
//	}
//	mid := randomizedPartion(nums, l, r)
//	randomizedQuicksort(nums, l, mid - 1)
//	randomizedQuicksort(nums, mid + 1, r)
//}

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

func quickSort(arr []int, start, end int) []int {
	if start < end {
		mid := arr[start]
		i := start
		j := end

		for i < j {

			for arr[j] >= mid && i < j {
				j--
			}
			for arr[i] <= mid && i < j {
				i++
			}
			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		arr[start] = arr[i]
		arr[i] = mid

		//fmt.Println(arr)
		quickSort(arr, start, j-1)
		quickSort(arr, j+1, end)
	}
	return arr
}

//func quickSort(arr[]int, start, end int) []int {
//	if start < end {
//		mid := start
//		i, j := start, end
//		for i < j {
//			for arr[j] >= mid && i < j {
//				j--
//			}
//			for arr[i] < mid && i < j {
//				i++
//			}
//			if i < j {
//				arr[j], arr[i] = arr[i], arr[j]
//			}
//		}
//		arr[start] = arr[i]
//		arr[i] = mid
//
//		quickSort(arr, start, j - 1)
//		quickSort(arr, j + 1, end)
//	}
//	return arr
//}
