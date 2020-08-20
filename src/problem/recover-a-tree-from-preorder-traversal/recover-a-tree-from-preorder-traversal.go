package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(recoverFromPreorder("1-2--3--4-5--6--7"))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(S string) *TreeNode {
	nums_ := 0
	index := 0
	stackTreeNode := make([]*TreeNode, 0)
	var root *TreeNode
	for index < len(S) {
		val := 0
		for ; index < len(S) && S[index] == '-'; index++ {
			nums_++
		}
		for ; index < len(S) && S[index] >= '0' && S[index] <= '9'; index++ {
			numInt, _ := strconv.Atoi(string(S[index]))
			val = val*10 + numInt

		}
		if nums_ == 0 {
			root = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			stackTreeNodePush(&stackTreeNode, root)
		} else {
			if len(stackTreeNode) == nums_ { // 证明是左子
				fatherNode := stackTreeNodeGetTop(&stackTreeNode)
				leftTreeNode := &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
				(*fatherNode).Left = leftTreeNode
				stackTreeNodePush(&stackTreeNode, leftTreeNode)
			} else {
				stackTreeNodePop(&stackTreeNode, len(stackTreeNode)-nums_)
				fatherNode := stackTreeNodeGetTop(&stackTreeNode)
				rightTreeNode := &TreeNode{
					Val:   val,
					Left:  nil,
					Right: nil,
				}
				(*fatherNode).Right = rightTreeNode
				stackTreeNodePush(&stackTreeNode, rightTreeNode)

			}
			nums_ = 0
		}
	}
	return root
}

func stackTreeNodePush(treeNodeStack *[]*TreeNode, node *TreeNode) {
	*treeNodeStack = append(*treeNodeStack, node)
}

func stackTreeNodePop(treeNodeStack *[]*TreeNode, popNum int) {
	*treeNodeStack = (*treeNodeStack)[:len(*treeNodeStack)-popNum]
}

func stackTreeNodeGetTop(treeNodeStack *[]*TreeNode) **TreeNode {
	return &(*treeNodeStack)[len(*treeNodeStack)-1]
}
