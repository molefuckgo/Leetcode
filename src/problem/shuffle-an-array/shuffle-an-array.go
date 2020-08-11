package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{1, 2, 3}
	oldNums := make([]int, len(nums))
	copy(oldNums, nums)
	//obj := Constructor(nums)
	//obj.Reset()
	//obj.Shuffle()

}

func init() {
	rand.Seed(time.Now().Unix())
}

type Solution struct {
	nums, oldNums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums:    nums,
		oldNums: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	this.nums = this.oldNums
	fmt.Println("充值前old", this.oldNums)
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := len(this.nums) - 1; i >= 0; i-- {
		randNum := rand.Intn(i + 1)
		fmt.Println(randNum)
		this.nums[i], this.nums[randNum] = this.nums[randNum], this.nums[i]
	}
	return this.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
