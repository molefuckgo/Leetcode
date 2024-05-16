package main

import "fmt"

func main() {
	bst := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	var result = []int{}
	printTree(bst, &result)
	fmt.Println(result)
}

func printTree(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	printTree(root.Left, result)
	*result = append(*result, root.Val)
	printTree(root.Right, result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			nums[0],
			nil, nil,
		}
	}
	if len(nums) == 2 {
		return &TreeNode{
			nums[1],
			&TreeNode{
				nums[0],
				nil, nil,
			}, nil,
		}
	}

	return &TreeNode{
		nums[len(nums)/2],
		sortedArrayToBST(nums[:len(nums)/2]),
		sortedArrayToBST(nums[len(nums)/2+1:]),
	}
}
