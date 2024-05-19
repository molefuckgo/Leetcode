package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result = [][]int{}
	if root == nil {
		return result
	}
	var stack = []*TreeNode{root}

	for len(stack) > 0 {
		var curNodeNums = len(stack)
		var curResult = []int{}
		for i := 0; i < curNodeNums; i++ {

			var node = stack[0]
			stack = stack[1:]
			curResult = append(curResult, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		result = append(result, curResult)
	}
	return result
}
