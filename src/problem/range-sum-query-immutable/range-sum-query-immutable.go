package main

func main() {

}

type NumArray struct {
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
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
