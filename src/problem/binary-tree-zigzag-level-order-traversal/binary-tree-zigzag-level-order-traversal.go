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
	zigzagLevelOrder(root)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result = [][]int{}
	if root == nil {
		return result
	}
	var zigzag = false
	var queue = []*TreeNode{root}
	var curLaver = -1
	for len(queue) > 0 {
		var curNums = len(queue)
		zigzag = !zigzag
		curLaver++
		for i := 0; i < curNums; i++ {
			var node = queue[0]
			queue = queue[1:]
			if i == 0 {
				result = append(result, []int{})
			}
			result[curLaver] = append(result[curLaver], node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if curLaver%2 == 1 {
			for i, j := 0, len(result[curLaver])-1; i < j; i, j = i+1, j-1 {
				result[curLaver][i], result[curLaver][j] = result[curLaver][j], result[curLaver][i]
			}
		}
	}
	return result
}
