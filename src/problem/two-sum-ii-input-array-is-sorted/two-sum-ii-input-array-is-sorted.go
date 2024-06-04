package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}

func twoSum(numbers []int, target int) []int {
	var slow, fast = 0, len(numbers) - 1
	for numbers[slow]+numbers[fast] != target {
		if numbers[slow]+numbers[fast] < target {
			slow++
		} else {
			fast--
		}
	}
	return []int{slow + 1, fast + 1}
}
