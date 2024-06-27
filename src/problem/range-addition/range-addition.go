package main

import "fmt"

func main() {
	fmt.Println(getModifiedArray(5, [][]int{{
		1, 3, 2,
	},
		{
			2, 4, 3,
		},
		{
			0, 2, -2,
		}}))
}

type Difference struct {
	diff []int
}

func newDifference(nums []int) *Difference {
	var diff = make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &Difference{diff: diff}
}

func (d *Difference) Increment(i, j, val int) {
	d.diff[i] += val
	if j+1 < len(d.diff) {
		d.diff[j+1] -= val
	}
}

func (d *Difference) Result() []int {
	var result = make([]int, len(d.diff))
	result[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		result[i] = result[i-1] + d.diff[i]
	}
	return result
}

func getModifiedArray(length int, updates [][]int) []int {
	var nums = make([]int, length)

	var df = newDifference(nums)
	for _, update := range updates {
		df.Increment(update[0], update[1], update[2])
	}
	return df.Result()
}
