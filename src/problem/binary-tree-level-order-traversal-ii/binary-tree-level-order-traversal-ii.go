package main

func main() {
	/*
			1
		  / \
		 2   3
		/     \
		4       5
	*/
	var root = &TreeNode{
		1, nil, nil,
	}
	root.Left = &TreeNode{
		2, nil, nil,
	}
	root.Right = &TreeNode{
		3, nil, nil,
	}

	root.Left.Left = &TreeNode{
		4, nil, nil,
	}
	root.Right.Right = &TreeNode{
		5, nil, nil,
	}
	levelOrderBottom(root)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue = []*TreeNode{root}
	var curLayer = -1
	for len(queue) > 0 {
		curLayer++
		var curNums = len(queue)
		for i := 0; i < curNums; i++ {
			if i == 0 {
				result = append(result, []int{})
			}
			var node = queue[0]
			queue = queue[1:]
			result[curLayer] = append(result[curLayer], node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	reverseArray(result)
	return result
}

func reverseArray(arr [][]int) {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
