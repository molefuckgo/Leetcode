package main

import "fmt"

func main() {
	var threeNode = &ListNode{
		Val:  3,
		Next: nil,
	}
	var twoNode = &ListNode{
		Val:  2,
		Next: nil,
	}
	threeNode.Next = twoNode
	var zeroNode = &ListNode{
		Val:  0,
		Next: nil,
	}
	twoNode.Next = zeroNode
	var negativeFour = &ListNode{
		Val:  -4,
		Next: twoNode,
	}
	zeroNode.Next = negativeFour
	fmt.Println(detectCycle(threeNode).Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	var slow, fast = head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			slow = head
			for slow != fast {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
