package main

import "fmt"

func main() {
	headA := ListNode{
		Val:  3,
		Next: nil,
	}
	headB := ListNode{
		Val:  2,
		Next: &headA,
	}
	fmt.Println(getIntersectionNode(&headA, &headB))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curNodeA, curNodeB := headA, headB
	flag := 0
	if curNodeA == nil || curNodeB == nil {
		return nil
	}
	for curNodeA != curNodeB {
		if flag > 1 {
			return nil
		}
		if curNodeA == nil {
			curNodeA = headB
		}
		if curNodeB == nil {
			curNodeB = headA
		}
		curNodeA = curNodeA.Next
		if curNodeA == nil {
			flag++
			curNodeA = headB
		}
		curNodeB = curNodeB.Next
		if curNodeB == nil {
			curNodeB = headA
		}
	}
	return curNodeA
}
