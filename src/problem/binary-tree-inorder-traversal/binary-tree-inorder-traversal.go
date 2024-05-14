package main

import "fmt"

func main() {
	var root = TreeNode{
		1, nil, &TreeNode{
			2, &TreeNode{
				3, nil, nil,
			}, nil,
		},
	}
	fmt.Println(inorderTraversal(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func inorderTraversal(root *TreeNode) []int {
	var result = make([]int, 0)
	inorderTraversal2(root, &result)
	return result
}

func inorderTraversal2(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inorderTraversal2(root.Left, result)
	*result = append(*result, root.Val)
	inorderTraversal2(root.Right, result)
}
*/

func inorderTraversal(root *TreeNode) []int {
	var result = make([]int, 0)
	var stack = []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right
	}

	return result
}
