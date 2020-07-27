package main

import "fmt"

func main() {
	tree := TreeNode{0, nil, nil}
	fmt.Println(1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

}
