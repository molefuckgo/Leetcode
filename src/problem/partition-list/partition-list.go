package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	var dummy1 = &ListNode{
		Val: -1,
	}
	var dummy2 = &ListNode{
		Val: -1,
	}

	var p1, p2 = dummy1, dummy2
	var p = head
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		var temp = p.Next
		p.Next = nil
		p = temp
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
