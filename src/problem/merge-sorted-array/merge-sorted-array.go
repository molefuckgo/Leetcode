package main

import "fmt"

func main() {
	var nums1 = []int{2, 0}
	var nums2 = []int{1}

	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	if n == 0 {
		return
	}

	realNums1 := make([]int, m)
	copy(realNums1, nums1[:m])

	var nums1Index, nums2Index = 0, 0
	var i = 0
	for ; nums1Index < m && nums2Index < n; i++ {
		if realNums1[nums1Index] <= nums2[nums2Index] {
			nums1[i] = realNums1[nums1Index]
			nums1Index++
		} else {
			nums1[i] = nums2[nums2Index]
			nums2Index++
		}
	}
	if nums1Index < m {
		copy(nums1[i:], realNums1[nums1Index:])
	} else if nums2Index < n {
		copy(nums1[i:], nums2[nums2Index:])
	}
}
