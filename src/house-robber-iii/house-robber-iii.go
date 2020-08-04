package main

import "fmt"

func main() {
	root := TreeNode{
		Val: 3,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	root.Left.Right = &TreeNode{
		Val: 3,
	}
	root.Right.Right = &TreeNode{
		Val: 1,
	}
	fmt.Println(rob(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	val := dfs(root)
	return max(val[0], val[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	l, r := dfs(node.Left), dfs(node.Right)
	selected := node.Val + l[1] + r[1]
	notSelected := max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
