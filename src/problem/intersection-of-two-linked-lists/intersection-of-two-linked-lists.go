package main

import "fmt"

func main() {
	// 创建相交节点及其后续节点
	intersectNode := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}

	// 创建第一个链表：4 -> 1 -> (8 -> 4 -> 5)
	headA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: intersectNode}}

	// 创建第二个链表：5 -> 6 -> 1 -> (8 -> 4 -> 5)
	headB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: intersectNode}}}
	fmt.Println(getIntersectionNode(headA, headB).Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var a, b = headA, headB
	for a != b && a != nil {
		if a.Next == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b.Next == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
