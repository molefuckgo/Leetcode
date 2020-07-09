package main

import "fmt"

func main() {
	// var nums1 = []int{1, 3, 9, 10}
	// var nums2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var nums1 = []int{1, 1, 1, 1}
	// var nums2 = []int{0, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var nums1 = []int{1, 2}
	// var nums2 = []int{3, 4}
	// var nums1 = []int{}
	// var nums2 = []int{1, 2, 3, 4}
	// var nums1 = []int{}
	// var nums2 = []int{1}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0 // 切除数组的长度
	for {
		if index1 == len(nums1) { // 第一个数组完全切没了
			return nums2[index2+k-1] //
		}
		if index2 == len(nums2) { // 第二个数组完全切没了
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
