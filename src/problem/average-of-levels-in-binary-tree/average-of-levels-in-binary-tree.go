package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{0}
	}
	var result = []float64{}
	var stack = []*TreeNode{root}

	for len(stack) > 0 {
		var curLen = len(stack)
		var curSum float64
		for i := 0; i < curLen; i++ {
			var node = stack[0]
			stack = stack[1:]
			curSum += float64(node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		result = append(result, curSum/float64(curLen))
	}
	return result
}
