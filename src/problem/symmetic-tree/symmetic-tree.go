package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create a sample tree:
	//      1
	//     / \
	//    2   3
	//   / \
	//  4   5
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println(isSymmetric(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}
*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	result := levelOrder(root)
	for i := 1; i < len(result); i++ {
		if len(result[i])%2 != 0 {
			return false
		}
		var left = result[i][:len(result[i])/2]
		var right = result[i][len(result[i])/2:]
		reverseSlice(right)
		for i := 0; i < len(left); i++ {
			if left[i] != right[i] {
				return false
			}
		}
	}
	return true
}

func reverseSlice(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return i > j
	})
}

func levelOrder(root *TreeNode) (result [][]int) {
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var levelSize = len(queue)
		var curLevelResult []int
		for i := 0; i < levelSize; i++ {
			var node = queue[0]
			queue = queue[1:]
			if node == nil {
				curLevelResult = append(curLevelResult, -10000)
			} else {
				curLevelResult = append(curLevelResult, node.Val)
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
		result = append(result, curLevelResult)
	}
	return result
}
