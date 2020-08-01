package main

func main() {

}

func singleNumbers(nums []int) []int {
	numsLen := len(nums)
	result := make([]int, 0)
	numsInfo := make(map[int]int)

	for i := 0; i < numsLen; i++ {
		numsInfo[nums[i]] += 1

		// if numsInfo[nums[i]] == 2 {
		// 	delete(numsInfo, nums[i])
		// }
	}
	for k, v := range numsInfo {
		if v == 1 {
			result = append(result, k)
		}
	}
	return result
}
