package main

func main() {

}

/*func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}*/

func containsDuplicate(nums []int) bool {
	var numsHash = make(map[int]bool)

	for _, num := range nums {
		if numsHash[num] == true {
			return true
		}
		numsHash[num] = true
	}
	return false
}
