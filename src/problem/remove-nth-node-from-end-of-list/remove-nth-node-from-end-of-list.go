package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{
		Val: 1,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next = &ListNode{
		Val: 3,
	}
	head.Next.Next.Next = &ListNode{
		Val: 4,
	}
	head.Next.Next.Next.Next = &ListNode{
		Val: 5,
	}
	fmt.Println(removeNthFromEnd(&head, 2))
}

func removeNth(head *ListNode, n int) *ListNode {
	var a = head
	for i := 1; i < n-1; i++ {
		a = a.Next
	}
	a.Next = a.Next.Next
	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy = &ListNode{
		Val:  0,
		Next: head,
	}
	var first, second = dummy, dummy
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

/*func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	first := dummy
	second := dummy
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}*/
