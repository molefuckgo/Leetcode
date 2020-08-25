package main

func main() {

}
func exchange(nums []int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left]&1 == 0 && nums[right]&1 == 1 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		} else if nums[left]&1 == 0 && nums[right]&1 == 0 {
			right--
		} else if nums[left]&1 == 1 && nums[right]&1 == 1 {
			left++
		} else {
			left++
			right--
		}
	}
	return nums
}
