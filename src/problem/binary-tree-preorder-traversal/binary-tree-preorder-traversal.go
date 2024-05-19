package main

import "fmt"

func main() {
	fmt.Println(preorderTraversal(&TreeNode{
		1, nil, nil,
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func preorderTraversal(root *TreeNode) []int {
	var result = make([]int, 0)
	var trav func(root *TreeNode)
	trav = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		trav(root.Left)
		trav(root.Right)
	}
	trav(root)
	return result
}
*/
// 二叉树前序遍历
func preorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int
	if root == nil {
		return result
	}

	stack = append(stack, root)
	for len(stack) != 0 {
		var node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}
