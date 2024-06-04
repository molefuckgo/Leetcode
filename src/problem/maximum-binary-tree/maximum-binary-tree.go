package main

import "fmt"

func main() {
	var result = constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}
	var max, maxIdx = -1, -1

	for i, num := range nums {
		if num > max {
			max = num
			maxIdx = i
		}
	}
	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:maxIdx]),
		Right: constructMaximumBinaryTree(nums[maxIdx+1:]),
	}
}
