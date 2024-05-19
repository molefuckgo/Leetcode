package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*func postorderTraversal(root *TreeNode) []int {
	var result []int

	var pt func(root *TreeNode)

	pt = func(root *TreeNode) {
		if root == nil {
			return
		}
		pt(root.Left)
		pt(root.Right)
		result = append(result, root.Val)
	}
	pt(root)
	return result
}*/
// 二叉树后序遍历
func postorderTraversal(root *TreeNode) []int {
	var result []int
	var stack = []*TreeNode{}
	var prev *TreeNode
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || prev == root.Right {
			result = append(result, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return result
}
