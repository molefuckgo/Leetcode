package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums, 3, 0, 0)
}

/**
nums:数组
n:计算几个数之和
start:从第几个数字开始计算
target:目标
*/
func nSumTarget(nums []int, n, start, target int) (res [][]int) {
	var sz = len(nums)
	if sz < n || n < 2 {
		return
	}
	if n == 2 {
		var low = start
		var high = sz - 1
		for low < high {
			var left = nums[low]
			var right = nums[high]
			var sum = left + right
			if sum > target {
				high--
			} else if sum < target {
				low++
			} else {
				res = append(res, []int{left, right})
				for low < high && left == nums[low] {
					low++
				}
				for high > low && right == nums[high] {
					high--
				}
			}
		}
	} else {
		for i := start; i < sz; i++ {
			var sub = nSumTarget(nums, n-1, i, target-nums[i])
			for _, arr := range sub {
				arr = append(arr, nums[i])
				res = append(res, arr)
			}
			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return
}
