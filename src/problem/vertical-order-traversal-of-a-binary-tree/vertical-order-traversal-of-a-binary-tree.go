package main

import (
	"math"
	"sort"
)

func main() {
	/* 3
	 / \
	9  20
	  /  \
	 15   7
	*/
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	verticalTraversal(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type data struct {
	col, row, val int
}

func verticalTraversal(root *TreeNode) [][]int {
	var result = [][]int{}
	var nodes = []data{}
	var dfs func(node *TreeNode, col, row int)

	dfs = func(node *TreeNode, col, row int) {
		if node == nil {
			return
		}
		nodes = append(nodes, data{col: col, row: row, val: node.Val})
		dfs(node.Left, col+1, row-1)
		dfs(node.Right, col+1, row+1)
	}
	dfs(root, 0, 0)

	sort.Slice(nodes, func(i, j int) bool {
		var a, b = nodes[i], nodes[j]
		return a.row < b.row || a.row == b.row && (a.col < b.col || a.col == b.col && a.val < b.val)
	})

	var minRow = math.MinInt32
	for _, node := range nodes {
		if minRow != node.row {
			minRow = node.row
			result = append(result, nil)
		}
		result[len(result)-1] = append(result[len(result)-1], node.val)
	}
	return result
}

/*
type data struct{ col, row, val int }

func verticalTraversal(root *TreeNode) (ans [][]int) {
	nodes := []data{}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		nodes = append(nodes, data{col, row, node.Val})
		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)

	sort.Slice(nodes, func(i, j int) bool {
		a, b := nodes[i], nodes[j]
		return a.col < b.col || a.col == b.col && (a.row < b.row || a.row == b.row && a.val < b.val)
	})

	lastCol := math.MinInt32
	for _, node := range nodes {
		if node.col != lastCol {
			lastCol = node.col
			ans = append(ans, nil)
		}
		ans[len(ans)-1] = append(ans[len(ans)-1], node.val)
	}
	return
}*/
