package main

import "fmt"

func main() {
	tree := TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	tree.Left = &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	//a := 1
	//fmt.Println(a)
	fmt.Println(maxDepth(&tree))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth2(root.Left), maxDepth2(root.Right)) + 1

}
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDeep := 0
	var queue = []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		thisQueueLen := len(queue)
		for ; thisQueueLen > 0; thisQueueLen-- {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		maxDeep++
	}
	return maxDeep
}
