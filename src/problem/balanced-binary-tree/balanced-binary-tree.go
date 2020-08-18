package main

import (
	"fmt"
	"math"
)

func main() {
	//[1,2,2,3,3,3,3,4,4,4,4,4,4,null,null,5,5]
	root := TreeNode{
		Val: 1,
	}

	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 2,
	}

	root.Left.Left = &TreeNode{
		Val: 3,
	}
	root.Left.Right = &TreeNode{
		Val: 3,
	}

	root.Left.Left.Left = &TreeNode{
		Val: 4,
	}
	root.Left.Left.Right = &TreeNode{
		Val: 4,
	}

	fmt.Println(isBalanced(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	// if root == nil || (root.Left == nil && root.Right == nil) {

}

func frontOrderPrint(node *TreeNode, curHeight, minHeight, maxHeight int) (int, int) {
	//fmt.Println("fuck", node == nil)

}
