package main

func main() {

}

func maxSubArray(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < numsLen; i++ {
		nums[i] = max2(nums[i], nums[i]+nums[i-1])
		max = max2(max, nums[i])
	}
	return max

}

func max2(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
