package main

func main() {
	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	deleteDuplicates(&head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	hair := ListNode{Val: 0}
	hair.Next = head // head是当前节点
	for head != nil && head.Next != nil {
		if head.Val != head.Next.Val {
			head = head.Next
		} else {
			nowPtr := head.Next
			for nowPtr != nil && nowPtr.Next != nil && nowPtr.Next.Val == nowPtr.Val {
				nowPtr = nowPtr.Next
			}
			head.Next = nowPtr.Next
		}
	}
	return hair.Next
}
