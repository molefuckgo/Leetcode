package main

import (
	"fmt"
	"math"
)

func main() {
	root := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
		},
		Right: nil,
	}
	fmt.Println(minDepth(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minHeight := math.MaxInt64
	if root.Left != nil {
		minHeight = min2(minDepth(root.Left), minHeight)
	}
	if root.Right != nil {
		minHeight = min2(minDepth(root.Right), minHeight)
	}
	return minHeight + 1
}

func min2(x, y int) int {
	if x >= y {
		return y
	}
	return x
}
