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
	//head.Next.Next = &ListNode{
	//	Val: 3,
	//}
	//head.Next.Next.Next = &ListNode{
	//	Val: 4,
	//}
	//head.Next.Next.Next.Next = &ListNode{
	//	Val: 5,
	//}
	fmt.Println(removeNthFromEnd(&head, 2))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
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
}

//public ListNode removeNthFromEnd(ListNode head, int n) {
//    ListNode dummy = new ListNode(0);
//    dummy.next = head;
//    ListNode first = dummy;
//    ListNode second = dummy;
//    // Advances first pointer so that the gap between first and second is n nodes apart
//    for (int i = 1; i <= n + 1; i++) {
//        first = first.next;
//    }
//    // Move first to the end, maintaining the gap
//    while (first != null) {
//        first = first.next;
//        second = second.next;
//    }
//    second.next = second.next.next;
//    return dummy.next;
//}
//
