package main

import "fmt"

func main() {
	head := ListNode{
		Val:  1,
		Next: nil,
	}
	head.Next = &ListNode{
		Val:  5,
		Next: nil,
	}
	h := reverseList(&head)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curHead := head.Next
	head.Next = nil
	for curHead != nil {
		curNext := curHead.Next
		curHead.Next = head
		head = curHead
		curHead = curNext
	}
	return head
}
