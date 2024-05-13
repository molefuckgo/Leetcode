package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
import (
	"fmt"
)

func main() {
	var nums = []int{0, 0, 1, 2, 2, 2, 4, 4}
	var expectedNums = []int{0, 1, 2, 4}
	k := removeDuplicates(nums)
	if k != len(expectedNums) {
		fmt.Println("Failed: Length does not match.")
		return
	}
	for i := 0; i < k; i++ {
		if nums[i] != expectedNums[i] {
			fmt.Printf("Failed: Element at index %d does not match.\n", i)
			return
		}
	}
	fmt.Println("Passed: removeDuplicates works correctly.")
}

//
//func removeDuplicates(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	var last = 0
//	for i := 1; i < len(nums); i++ {
//		if nums[i] != nums[last] {
//			last++
//			nums[last] = nums[i]
//		}
//	}
//	return last + 1
//}

//func removeDuplicates(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//
//	// last 指向最后一个不重复的元素，i 用于遍历数组。
//	last := 0
//	for i := 1; i < len(nums); i++ {
//		if nums[i] != nums[last] {
//			// 发现一个新的不重复元素，移动 last 指针并更新该位置的值。
//			last++
//			nums[last] = nums[i]
//		}
//	}
//
//	// 返回新切片的长度，即最后一个不重复元素的索引 + 1。
//	return last + 1
//}

func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	var lastNum int
	for i := 0; i < len(nums); {
		if i == 0 && len(nums) != 0 {
			lastNum = nums[i]
			i++
		} else {
			if nums[i] == lastNum {
				numsLen--
				deleteArrayIndex(&nums, i)
			} else {
				lastNum = nums[i]
				i++
			}
		}

	}
	return numsLen
}

func deleteArrayIndex(nums *[]int, index int) {
	*nums = append((*nums)[:index], (*nums)[index+1:]...)
}
