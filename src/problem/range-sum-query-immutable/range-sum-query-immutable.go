package main

import "fmt"

func main() {
	var numArray = Constructor([]int{-1})
	fmt.Println(numArray.SumRange(0, 0))
}

/*type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i >= 0 && j <= len(this.nums) {
		sum := 0
		for ; i <= j; i++ {
			sum += this.nums[i]
		}
		return sum
	} else {
		return 0
	}
}*/

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	var numsLen = len(nums)
	var preSum = make([]int, numsLen+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}
	return NumArray{
		preSum: preSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > right || left < 0 {
		return 0
	}
	return this.preSum[right+1] - this.preSum[left]
}
