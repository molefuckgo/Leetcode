package main

func main() {
	head := ListNode{
		Val: 0,
	}
	head.Next = &ListNode{
		Val: 1,
	}
	//head.Next.Next = &ListNode{
	//	Val: 2,
	//}
	//head.Next.Next.Next = &ListNode{
	//	Val: 3,
	//}
	//head.Next.Next.Next.Next = &ListNode{
	//	Val: 4,
	//}
	//head.Next.Next.Next.Next.Next = &ListNode{
	//	Val: 5,
	//}
	sortedListToBST(&head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{
			Val: head.Val,
		}
	}
	listNodeList := make([]*ListNode, 0)
	for head != nil {
		listNodeList = append(listNodeList, head)
		head = head.Next
	}
	allmidNodeIndex := len(listNodeList) / 2
	rootTreeNode := TreeNode{
		Val: listNodeList[allmidNodeIndex].Val,
	}
	leftNodeList := listNodeList[:allmidNodeIndex]
	rightNodeList := listNodeList[allmidNodeIndex+1:]
	rootTreeNode.Left = sortedListdp(leftNodeList)
	rootTreeNode.Right = sortedListdp(rightNodeList)

	return &rootTreeNode
	//	输入：
	//[0,1,2,3,4,5]
	//输出：
	//[3,2,4,1,null,null,5,0]
	//预期：
	//[3,1,5,0,2,4]
}

func sortedListdp(listListNode []*ListNode) *TreeNode {
	if len(listListNode) == 0 {
		return nil
	} else if len(listListNode) == 1 {
		return &TreeNode{
			Val: listListNode[0].Val,
		}
	} else if len(listListNode) == 2 {
		return &TreeNode{
			Val: listListNode[0].Val,
			Right: &TreeNode{
				Val: listListNode[1].Val,
			},
		}
	}
	allmidNodeIndex := len(listListNode) / 2
	leftResult := sortedListdp(listListNode[:allmidNodeIndex])
	rightResult := sortedListdp(listListNode[allmidNodeIndex+1:])
	rootTreeNode := &TreeNode{
		Val:   listListNode[allmidNodeIndex].Val,
		Left:  leftResult,
		Right: rightResult,
	}
	return rootTreeNode

}
